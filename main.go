package main

import (
	"fmt"
	"strconv"
)

func ScanPort(target string) {
	for port := 1; port <= 10; port++ {
		// protocol := "tcp"
		fmt.Println(strconv.Itoa(port))
	}
}

func main() {
	ScanPort("nil")
}
