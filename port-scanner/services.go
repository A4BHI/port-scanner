package portscanner

import (
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

	b := []byte{}

	_, err = file.Read(b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
	//fuvk
}
