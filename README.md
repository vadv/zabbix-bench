Benchmarking and stress testing tool for the Zabbix server

```bash
Usage of zabbix-bench:
  -client int
        number of concurrent clients (default 200)
  -client-format string
        format of client name (default "client-%d")
  -metric-format string
        format of metric name in packet (default "metric-%d")
  -packet-delay duration
        delay of send packet (default 100ms)
  -packet-send-timeout duration
        packet send timeout (default 10ms)
  -packet-size int
        count of metric in packet (default 100)
  -zabbix string
        address of zabbix server (default "127.0.0.1:10051")
```
