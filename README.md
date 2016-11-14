Benchmarking and stress testing tool for the Zabbix server

## Populate data

Run `./generate-template-metrics.sh > template.xml` and upload template with metrics.
Run `./generate-template-clients.sh > clients.xml` and upload template with clients.

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
