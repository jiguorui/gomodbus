package modbus

import (
	"errors"
)

//Modbus Protocol Data Unit(PDU)
type MbReqPDU struct {
	b []byte
}

func NewMbReqPDU(b []byte) (*MbReqPDU, error) {
	if len(b) != 5 {
		return nil, errors.New("length of bytes is not 5.")
	}
	m := &MbReqPDU{b}
	return m, nil
}

func (m *MbReqPDU) Decode() (byte, int16, int16) {
	return m.b[0], getInt16(m.b[1], m.b[2]), getInt16(m.b[3], m.b[4])
}

func (m *MbReqPDU) Encode(fcode byte, start, qty int16) {
	m.b[0] = fcode
	m.b[1], m.b[2] = getBytes(start)
	m.b[3], m.b[4] = getBytes(qty)
}
