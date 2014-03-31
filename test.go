package main

import (
	"fmt"
	"github.com/jiguorui/gomodbus/modbus"
)

func main() {
	var b [160]byte
	m, _ := modbus.NewMbADU(b[0:7])
	m.Encode(1, 0, 3, 1)
	tid, pid, length, unit := m.Decode()
	fmt.Printf("%d, %d, %d, %d\n", tid, pid, length, unit)
	fmt.Printf("%d, %d, %d, %d\n", b[1], b[3], b[5], b[6])

	m2, _ := modbus.NewMbReqPDU(b[7:12])
	m2.Encode(3, 0, 12)
	c, s, n := m2.Decode()
	fmt.Printf("%d, %d, %d\n", c, s, n)
	fmt.Printf("%d, %d, %d\n", b[7], b[9], b[11])

	srv, err := modbus.NewServer(":502")
	if err != nil {
		fmt.Println(err)
	}
	defer srv.Close()
	srv.Loop()
}
