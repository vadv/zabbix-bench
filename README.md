Benchmarking and stress testing tool for the Zabbix server

## Populate data

* Run `./generate-template-metrics.sh > template.xml` and upload template with metrics.

* Run `./generate-template-clients.sh > clients.xml` and upload template with clients.

## Choose options

```bash
./bin/zabbix-bench  -h
Usage of ./bin/zabbix-bench:
  -client int
        number of concurrent clients (default 200)
  -client-format string
        format of client name (default "client-%d")
  -max-duration duration
        max duration of benchmark test
  -max-metrics int
        max number of metrics sends
  -metric-format string
        format of metric name in packet (default "metric-%d")
  -packet-delay duration
        delay of send packet (default 100ms)
  -packet-send-timeout duration
        packet send timeout (default 10ms)
  -packet-size int
        count of metric in packet (default 400)
  -threads int
        number of threads (default 8)
  -zabbix string
        address of zabbix server (default "127.0.0.1:10051")
```

## Run benchmark

```bash
make && ./bin/zabbix-bench -zabbix "192.168.1.1:10051" -packet-delay 1ms
go build -o ./bin/zabbix-bench ./src/main.go
Start 200 clients with packet size 100 metric and delay between packets 1ms
progress 1 s, 13100 metric/s
progress 2 s, 15300 metric/s
progress 3 s, 11200 metric/s
progress 4 s, 17500 metric/s
progress 5 s, 13100 metric/s
^C
-----------------------------
Total processed: 70300 (11716 metric/s)
```
