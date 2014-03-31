package modbus

//b1,b2 byte -> i16 int16
func getInt16(b1, b2 byte) int16 {
	var i16 int16
	i16 = int16(b1)<<8 + int16(b2)
	return i16
}

//i16 int16 -> b1, b2 byte
func getBytes(i16 int16) (h byte, l byte) {
	l = byte(i16 & 0x00FF)
	h = byte(i16 >> 8)
	return h, l
}
