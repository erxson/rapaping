package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"tcp/rapaping"

	"github.com/gookit/color"
)

const (
	infoFormat = ` Address Details:
  │ Address:     [ %s ]
  │ Hostname:    [ %s ]
  │ City:        [ %s ]
  │ Region:      [ %s ]
  │ Country:     [ %s ]
  │ ASN/ISP:     [ %s ]

`
)

var logger = log.New(os.Stdout, "", 0)

func main() {
	if len(os.Args) != 3 {
		color.Error.Printf("Usage: go run main.go ip port")
	}

	host := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])
	if err != nil || !rapaping.IsValidPort(port) {
		color.Error.Printf("Invalid port number:", err)
	}

	logger.Println("rapaping v1.4.88 - Copyright (c) 1917 Big Armenian Daddy")
	logger.Printf("\nConnecting to "+color.Green.Sprint("%s")+" on "+color.Green.Sprint("TCP %d")+":\n\n", host, port)

	ips, err := net.LookupIP(host)
	if err != nil {
		color.Error.Printf("Failed to resolve %s: %v", host, err)
		os.Exit(1)
	} else {
		ipInfo, err := rapaping.GetIPInfo(ips[0].String())
		if err == nil {
			logger.Printf(infoFormat, color.Green.Sprint(strings.Join(rapaping.TransformIPArray(ips), ", ")), color.Green.Sprint(ipInfo.Hostname), color.Green.Sprint(ipInfo.City), color.Green.Sprint(ipInfo.Region), color.Green.Sprint(ipInfo.Country), color.Green.Sprint(ipInfo.Org))
		}
	}

	info := &rapaping.Info{
		Host: host,
		Port: port,
	}

	rapaping.Run(info)
}
