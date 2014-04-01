package modbus

import (
	"net"
)

type Server struct {
	addr string
	lis  net.Listener
}

func handleConn(c net.Conn) {
	var b1 [261]byte
	var b2 [261]byte

	defer c.Close()

	m, _ := NewConn(c, b1[0:], b2[0:])

	for {
		err := m.StepHandle()
		if err != nil {
			break
		}
	}

}

func NewServer(addr string) (*Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Server{addr, lis}, nil
}

func (s *Server) Close() error {
	return s.lis.Close()
}

func (s *Server) DoLoop() error {
	for {
		c, err := s.lis.Accept()
		if err != nil {
			return err
		}
		go handleConn(c)
	}
	return nil
}
