package portscanner

import (
	"net"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ScanPort(target string, updates *tgbotapi.Update) {
	var OpenPorts []string
	for port := 1; port <= 1024; port++ {

		go func() {
			conn, err := net.DialTimeout("tcp", target+":"+strconv.Itoa(port), 500*time.Millisecond)
			if err != nil {

				return
			}
			OpenPorts = append(OpenPorts, strconv.Itoa(port))

			conn.Close()
		}()

	}

	extraPorts := []string{"8080", "3389", "1443", "3306", "3389", "5900"}
	go func() {
		for _, ports := range extraPorts {
			address := target + ":" + ports
			conn, err := net.DialTimeout("tcp", address, 1000*time.Millisecond)
			if err != nil {
				continue
			}
			OpenPorts = append(OpenPorts, ports)
			conn.Close()
		}
	}()

}
