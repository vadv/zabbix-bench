package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
	"zabbix"
)

var (
	argClientCount = flag.Int("client", 200, "number of concurrent clients")
	argThreadCount = flag.Int("threads", runtime.NumCPU()*4, "number of threads")
	argClinetName  = flag.String("client-format", "client-%d", "format of client name")
	argPacketSize  = flag.Int("packet-size", 400, "count of metric in packet")
	argMetricName  = flag.String("metric-format", "metric-%d", "format of metric name in packet")
	argPacketDelay = flag.Duration("packet-delay", 100*time.Millisecond, "delay of send packet")
	argSendTimeout = flag.Duration("packet-send-timeout", 10*time.Millisecond, "packet send timeout")
	argZabbix      = flag.String("zabbix", "127.0.0.1:10051", "address of zabbix server")
	argMaxMetrics  = flag.Int("max-metrics", 0, "max number of metrics each client sends")
	argMaxDuration = flag.Duration("max-duration", 0, "max duration of benchmark test")

	errorChannel     = make(chan error, 10)
	completedChannel = make(chan int, 10)
	signalChannel    = make(chan os.Signal, 1)

	mutex               = &sync.Mutex{}
	counter, total, sec = 0, 0, 1
)

func main() {

	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, syscall.SIGTERM)

	if !flag.Parsed() {
		flag.Parse()
	}

	runtime.GOMAXPROCS(*argThreadCount)
	for i := 0; i < *argClientCount; i++ {
		go StartClient(i)
	}

	report_and_exit := func() {
		speed := 0
		if sec > 0 {
			speed = int(total / sec)
		}
		fmt.Printf("\n-----------------------------\n")
		fmt.Printf("Total processed: %d (%d metric/s)\n", total, speed)
		os.Exit(0)
	}

	fmt.Printf("Start %d clients with packet size %d metric and delay between packets %v\n", *argClientCount, *argPacketSize, *argPacketDelay)
	ticker := time.Tick(time.Second)
	for {
		select {
		case <-ticker:
			mutex.Lock()
			fmt.Printf("progress %d s, %d metric/s\n", sec, counter)
			sec += 1
			counter = 0
			if *argMaxMetrics > 0 && total > *argMaxMetrics {
				report_and_exit()
			}
			if *argMaxDuration > 0 && time.Duration(sec)*time.Second > *argMaxDuration {
				report_and_exit()
			}
			mutex.Unlock()
		case count := <-completedChannel:
			mutex.Lock()
			counter += count
			total += count
			mutex.Unlock()
		case err := <-errorChannel:
			fmt.Fprintf(os.Stderr, "Error write metric:\t%s\n", err.Error())
		case <-signalChannel:
			report_and_exit()
		}
	}

}

// client of zabbix server
type client struct {
	id     int
	host   string
	sender *zabbix.Sender
}

// generate and send zabbix packet
func (c *client) send() error {
	now := time.Now().Unix()
	metrics := make([]*zabbix.Metric, 0)
	for i := 0; i < *argPacketSize; i++ {
		metrics = append(metrics, zabbix.NewMetric(c.host, fmt.Sprintf(*argMetricName, i), fmt.Sprintf("%d", i), now))
	}
	return c.sender.Send(zabbix.NewPacket(metrics, now))
}

func StartClient(id int) {
	time.Sleep((time.Duration(rand.Float64()*100) * *argPacketDelay) / 100)
	c := &client{
		id:     id,
		host:   fmt.Sprintf(*argClinetName, id),
		sender: zabbix.NewSender(*argZabbix),
	}
	ticker := time.Tick(*argPacketDelay)
	for {
		select {
		case <-ticker:
			if err := c.send(); err != nil {
				errorChannel <- err
			} else {
				completedChannel <- *argPacketSize
			}
		}
	}
}
