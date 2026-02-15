package portscanner

import (
	"bufio"
	"os"
	"strings"
)

type ServicesAndProtocols struct {
	NameOfService string
	Protocol      string
}
type DB struct {
	Port map[string]ServicesAndProtocols
}

func (db *DB) LookUP(portno string) (serivcename string, protocol string) {
	if name, ok := db.Port[portno]; ok {
		return name.NameOfService, name.Protocol
	}

	return "unknown", "unknown"
}

func LoadService(path string) (*DB, error) {
	file, err := os.Open(path)
	if err != nil {
		return &DB{}, err
	}
	defer file.Close()
	db := &DB{
		Port: make(map[string]ServicesAndProtocols),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines := scanner.Text()
		if len(lines) == 0 {
			continue
		}
		if lines[0] == '#' {
			continue
		}

		field := strings.Fields(lines)
		ports := strings.Split(field[1], "/")

		if ports[1] != "tcp" {
			continue
		}
		db.Port[ports[0]] = ServicesAndProtocols{
			NameOfService: field[0],
			Protocol:      ports[1],
		}

	}

	return db, err
}
