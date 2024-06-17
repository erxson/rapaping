package rapaping

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/gookit/color"
)

func colorPing(ping float64) string {
	var colorFunc func(a ...any) string

	switch {
	case ping < 200:
		colorFunc = color.Green.Sprint
	case ping < 400:
		colorFunc = color.Yellow.Sprint
	case ping < 1000:
		colorFunc = color.Red.Sprint
	default:
		colorFunc = color.BgRed.Sprint
	}

	return colorFunc(fmt.Sprintf("%.2fms", ping))
}

func TransformIPArray(ipArray []net.IP) []string {
        s := make([]string,0)
    for _, ip := range ipArray {
        s = append(s, ip.String())
    }
    return s
}

func GetIPInfo(ip string) (*IPInfo, error) {
	resp, err := http.Get(fmt.Sprintf("http://ipinfo.io/%s/json", ip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ipInfo IPInfo
	err = json.NewDecoder(resp.Body).Decode(&ipInfo)
	if err != nil {
		return nil, err
	}

	return &ipInfo, nil
}

func IsValidPort(port int) bool {
	return port > 0 && port < 65536
}