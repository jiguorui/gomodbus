//Reference:
//[1] MODBUS Application Protocol Specification V1.1b3.pdf
//[2] MODBUS Messaging on TCP/IP Implementation Guide V1.0b

package main

import (
	//	"errors"
	"fmt"
	"github.com/jiguorui/gomodbus/modbus"
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

// func readBytes(c net.Conn, b []byte) (int, error) {
// 	end := len(b)
// 	start := 0
// R:
// 	n, err := c.Read(b[start:end])
// 	if err != nil {
// 		return n, err
// 	}
// 	start = start + n
// 	if start < end {
// 		goto R
// 	}
// 	return end, nil
// }

// func procReq(req_pdu, rsp_pdu []byte) (byte, error) {
// 	funccode := req_pdu[0]
// 	if funccode == 3 {
// 		rsp_pdu[0] = funccode
// 		//startAddr := rsp_pdu[1] * 0x80 + rsp_pdu[2]
// 		qtyR := req_pdu[3]*0x80 + req_pdu[4]
// 		rsp_pdu[1] = qtyR * 2
// 		return qtyR*2 + 2, nil
// 	}
// 	rsp_pdu[0] = 0x80 + funccode
// 	rsp_pdu[1] = 1
// 	return 2, errors.New("Invalid function code")
// }

func handleConn(c net.Conn) {
	//see [1], max size of ADU is 260?
	//indication(request of client side) and response Application Data Unit(ADU)
	var b1 [261]byte
	var b2 [261]byte

	defer c.Close()

	m, _ := modbus.NewModbus(c, b1[0:], b2[0:])

	for {
		err := m.StepHandle()
		if err != nil {
			break
		}
	}

	// for {
	// 	//read MBAP header, see [1]
	// 	_, err := readBytes(c, indADU[0:7])
	// 	if checkError(err) {
	// 		break
	// 	}
	// 	req err := NewMbADU(indADU[0:7])
	// 	if checkError(err) {
	// 		break
	// 	}
	// 	tId, pId, length, unit := req.Decode()

	// 	_, err = readBytes(c, indADU[7:length + 7 - 1])
	// 	if checkError(err) {
	// 		break
	// 	}

	// 	n, err := procReq(indADU[7:dataLen], rspADU[7:])
	// 	//set MBAP header's 'length' byte
	// 	rspADU[5] = n + 1
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	//err or not, just write rsp to client
	// 	_, err = c.Write(rspADU[0 : n+7])
	// 	if checkError(err) {
	// 		break
	// 	}
	// }
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
