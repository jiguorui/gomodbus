package modbus

import (
	"errors"
)

//Modbus Protocol Data Unit(PDU)
type PDU struct {
	b []byte
}

func NewPDU(b []byte) (*PDU, error) {
	if len(b) < 5 {
		return nil, errors.New("length of bytes is less than 5.")
	}
	m := &PDU{b}
	return m, nil
}

func (m *PDU) Decode() (byte, int16, int16) {
	return m.b[0], getInt16(m.b[1], m.b[2]), getInt16(m.b[3], m.b[4])
}

func (m *PDU) Encode(fcode byte, start, qty int16) {
	m.b[0] = fcode
	m.b[1], m.b[2] = getBytes(start)
	m.b[3], m.b[4] = getBytes(qty)
}
