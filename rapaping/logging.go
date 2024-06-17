package rapaping

import (
	"log"
	"os"

	"github.com/gookit/color"
)

var logger = log.New(os.Stdout, "", 0)

func printReport(stats *ConnectionStats) {
	success := float64(stats.Connected) / float64(stats.Attempted) * 100
	logger.Printf("\nConnection statistics:\n")
	logger.Printf("        Attempted = "+color.Cyan.Sprint("%d")+", Connected = "+color.Green.Sprint("%d")+", Failed = "+color.Red.Sprint("%d")+" ("+color.Red.Sprint("%.2f%%")+")\n", stats.Attempted, stats.Connected, stats.Failed, success)
	logger.Printf("Approximate connection times:\n")

	if stats.Connected > 0 {
		averageTime := float64(stats.TotalTime.Milliseconds()) / float64(stats.Connected)
		logger.Printf("        Minimum = " + colorPing(float64(stats.MinTime.Milliseconds())) + ", Maximum = " + colorPing(float64(stats.MaxTime.Milliseconds())) + ", Average = " + colorPing(averageTime) + "\n\n")
	}
}
