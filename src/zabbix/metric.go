package zabbix

import (
	"time"
)

type Metric struct {
	Host  string `json:"host"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Clock int64  `json:"clock"`
}

func NewMetric(host, key, value string, clock ...int64) *Metric {
	m := &Metric{Host: host, Key: key, Value: value}
	if len(clock) > 0 {
		m.Clock = clock[0]
	} else {
		m.Clock = time.Now().Unix()
	}
	return m
}
