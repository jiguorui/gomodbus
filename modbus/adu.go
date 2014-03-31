package modbus

import (
	"errors"
)

//Modbus Applicatin Data Unit(ADU)
type MbADU struct {
	b []byte
}

func NewMbADU(b []byte) (*MbADU, error) {
	if len(b) != 7 {
		return nil, errors.New("length of bytes is not 7.")
	}

	m := &MbADU{b}
	return m, nil
}

func (m *MbADU) Decode() (int16, int16, int16, byte) {
	return getInt16(m.b[0], m.b[1]), getInt16(m.b[2], m.b[3]), getInt16(m.b[4], m.b[5]), m.b[6]
}

func (m *MbADU) Encode(tid, pid, length int16, unit byte) {
	m.b[0], m.b[1] = getBytes(tid)
	m.b[2], m.b[3] = getBytes(pid)
	m.b[4], m.b[5] = getBytes(length)
	m.b[6] = unit
}
