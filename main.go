package main

import (
	//	"fmt"
	"log"
	//	"net"
	"github.com/jiguorui/gomodbus/modbus"
)

func checkError(err error) bool {
	if err == nil {
		return false
	}
	log.Println(err)
	return true
}

func main() {
	srv, err := modbus.NewServer(":502")
	if checkError(err) {
		return
	}
	defer srv.Close()

	srv.DoLoop()
}
