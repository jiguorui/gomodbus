package modbus

import (
	"errors"
	"net"
)

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

type MbConn struct {
	c              net.Conn
	b1, b2         []byte
	reqADU, rspADU *MbADU
	reqPDU         *MbReqPDU
}

func NewMbConn(c net.Conn, b1, b2 []byte) (*MbConn, error) {
	if len(b1) < 260 || len(b2) < 260 {
		return nil, errors.New("buffer size too small")
	}

	reqADU, err := NewMbADU(b1[0:7])
	if err != nil {
		return nil, err
	}
	rspADU, err := NewMbADU(b2[0:7])
	if err != nil {
		return nil, err
	}
	reqPDU, err := NewMbReqPDU(b1[7:12])
	if err != nil {
		return nil, err
	}

	m := &MbConn{c, b1, b2, reqADU, rspADU, reqPDU}
	return m, nil
}

func (m *MbConn) StepHandle() error {
	_, err := readBytes(m.c, m.b1[0:7])
	if err != nil {
		return err
	}
	transID, protoID, length, unitID := m.reqADU.Decode()
	_, err = readBytes(m.c, m.b1[7:length+7-1])
	if err != nil {
		return err
	}
	fcode, _, qty := m.reqPDU.Decode()
	length = qty*2 + 3
	m.rspADU.Encode(transID, protoID, length, unitID)
	m.b2[7] = fcode
	m.b2[8] = byte(qty)
	_, err = m.c.Write(m.b2[0 : length+6])
	if err != nil {
		return err
	}
	return nil
}