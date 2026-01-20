package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ScanPort(target string, port int) {

	address := target + ":" + strconv.Itoa(port)

	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		fmt.Println("Closed Port:", port)
	}

	fmt.Println("Open port:", port)
	conn.Close()

}

func main() {
	for port := 1; port <= 1024; port++ {
		go ScanPort("", port)
		continue

	}
}
