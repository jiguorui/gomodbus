//Reference:
//[1] MODBUS Application Protocol Specification V1.1b3.pdf
//[2] MODBUS Messaging on TCP/IP Implementation Guide V1.0b

package main

import (
	"fmt"
	"log"
	"net"
)

func checkError(err error) bool {
	if err == nil {
		return false
	}
	log.Println(err)
	return true
}

func readBytes(c net.Conn, b []byte) (int, error) {
	end := len(b)
	start := 0
R:
	n, err := c.Read(b[start:end])
	if err != nil {
		return n, err
	}
	start = start + n
	if start < end {
		goto R
	}
	return end, nil
}

func handleConn(c net.Conn) {
	//see [1], max size of PDU is 260?
	var req_pdu [261]byte
	var rsp_pdu [261]byte

	defer c.Close()

	for {
		//read MBAP header, see [1]
		_, err := readBytes(c, req_pdu[0:7])
		if checkError(err) {
			break
		}
		//read data
		dataLen := req_pdu[5]
		_, err = readBytes(c, req_pdu[7:7+dataLen-1])
		if checkError(err) {
			break
		}

		funccode := req_pdu[7]
		//response
		copy(rsp_pdu[0:7], req_pdu[0:7])

		rsp_pdu[5] = 3
		rsp_pdu[7] = 0x83
		rsp_pdu[8] = 3
		_, err = c.Write(rsp_pdu[0:9])
		if checkError(err) {
			break
		}
	}
}

func main() {
	srv, err := net.Listen("tcp", ":502")
	if checkError(err) {
		return
	}
	defer srv.Close()
	for {
		conn, err := srv.Accept()
		fmt.Printf("connected\n")
		if checkError(err) {
			continue
		}
		go handleConn(conn)
	}
}
