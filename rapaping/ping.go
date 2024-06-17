package rapaping

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gookit/color"
)

// var dialer proxy.Dialer

func Run(info *Info) {
	stats := &ConnectionStats{}

	// Interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		printReport(stats)
		os.Exit(0)
	}()

	// Proxy Dialer
	// d, err := proxy.SOCKS5("tcp", info.Proxy, nil, &net.Dialer{Timeout: 5 * time.Second})
	// if err != nil {
	// 	color.Error.Printf("proxy dead mazafaka kurwa: %v", err)
	// }
	// dialer = d

	// Ping Loop
	for {
		ping(info.Host, info.Port, stats)
		time.Sleep(time.Millisecond * 1000)
	}
}

func ping(host string, port int, stats *ConnectionStats) {
	// Resolve and roll DNS
	ips, err := net.LookupIP(host)
	if err != nil {
		color.Error.Printf("Failed to resolve DNS for %s: %v\n", host, err)
		stats.Failed++
		return
	}

	rand.Seed(time.Now().UnixNano())
	ip := ips[rand.Intn(len(ips))]

	// Timer
	startTime := time.Now() // Start timer
	// conn, err := dialer.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 5*time.Second)
	if err != nil {
		color.Error.Println("Connection timed out")
		stats.Failed++
		return
	}
	defer conn.Close()
	duration := time.Since(startTime) // Stop timer

	logger.Printf("Connected to "+color.Green.Sprint("%s")+": time="+colorPing(float64(duration.Milliseconds()))+" protocol="+color.Green.Sprint("TCP")+" port="+color.Green.Sprint("%d")+"\n", ip, port)
	//logger.Printf("Connected to "+color.Green.Sprint("%s")+": time="+color.Green.Sprint("%.2fms")+ "\n", ip, float64(duration.Milliseconds()))

	// Stats
	stats.Connected++
	stats.TotalTime += duration

	if stats.MinTime == 0 || duration < stats.MinTime {
		stats.MinTime = duration
	}
	if duration > stats.MaxTime {
		stats.MaxTime = duration
	}
	stats.Attempted++
}
