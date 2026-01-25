package portscanner

import (
	"bufio"
	"fmt"
	"os"
)

type service struct {
	service string
	port    string
}

var slicesOfservice []service

func Services() {

	file, err := os.Open("/etc/services")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	//fuvk
}
