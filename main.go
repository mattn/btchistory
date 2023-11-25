package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type input struct {
	time time.Time
	in   float64
}

type rate struct {
	Pair string    `json:"pair"`
	Time time.Time `json:"time"`
	Rate string    `json:"rate"`
}

func timedRate(pair string, t time.Time) (float64, error) {
	u, err := url.Parse(`https://coincheck.com/ja/exchange/rates/search`)
	if err != nil {
		return 0, err
	}
	param := url.Values{}
	param.Set("pair", pair)
	param.Set("time", t.Format(`2006-01-02T15:04:05.000Z`))
	u.RawQuery = param.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var r rate
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(r.Rate, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func main() {
	var fname string
	var pair string
	var span int
	flag.StringVar(&fname, "input", "input.csv", "input CSV filename")
	flag.IntVar(&span, "span", 0, "interval of days across from start to end (0 is auto)")
	flag.StringVar(&pair, "pair", "btc_jpy", "pair name")
	flag.Parse()

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var inputs []input
	reader := csv.NewReader(f)
	for {
		cells, err := reader.Read()
		if err != nil {
			break
		}
		if len(cells) < 2 {
			continue
		}
		t, err := time.Parse(`2006/01/02 15:04:05`, cells[0])
		if err != nil {
			log.Fatal(err)
		}
		f, err := strconv.ParseFloat(cells[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, input{
			time: t,
			in:   f,
		})
	}
	if len(inputs) == 0 {
		log.Fatal("empty input.csv")
	}

	y, m, d := inputs[0].time.Date()
	curr := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	if span <= 0 {
		if len(inputs) >= 2 {
			y, m, d := inputs[len(inputs)-1].time.Date()
			end := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
			span = int(end.Sub(curr).Hours()/24) / 30
			if span < 1 {
				span = 1
			}
		} else {
			span = 1
		}
	}
	curr = curr.AddDate(0, 0, int(span))
	prev := curr

	total := 0.0
	for {
		curr = curr.AddDate(0, 0, int(span))
		over := 0
		for _, i := range inputs {
			if i.time.Before(curr) {
				if prev.Before(i.time) {
					total += i.in
				}
				over++
			}
		}
		r, err := timedRate(pair, curr)
		if err != nil {
			break
		}
		fmt.Printf("%s\t%f\t%f\n", curr.Format(`2006-01-02`), r, total*r)
		prev = curr
		if time.Now().Before(curr) {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}
