# btchistory



## Usage

Make input.csv from your BTC input history like below
```
2023/05/20 11:05:02,0.01
2023/07/01 04:07:54,0.05
2023/10/14 23:52:47,0.15
```

```
$ btchistory -input input.csv -span 10
2023-06-09	3686271.000000	0.000000
2023-06-19	3733234.000000	0.000000
2023-06-29	4348324.000000	0.000000
2023-07-09	4322806.000000	216140.300000
2023-07-19	4149704.000000	207485.200000
2023-07-29	4138005.000000	206900.250000
2023-08-08	4163000.000000	208150.000000
2023-08-18	3920006.000000	196000.300000
2023-08-28	3827795.000000	191389.750000
2023-09-07	3811170.000000	190558.500000
2023-09-17	3935829.000000	196791.450000
2023-09-27	3912759.000000	195637.950000
2023-10-07	4170600.000000	208530.000000
2023-10-17	4261376.000000	852275.200000
2023-10-27	5140448.000000	1028089.600000
2023-11-06	5234995.000000	1046999.000000
2023-11-16	5730812.000000	1146162.400000
```

![](screenshot.png)

## Installation

```
go install github.com/mattn/btchistory@latest
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)