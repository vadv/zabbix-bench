package zabbix

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

var zbxHeader = []byte("ZBXD\x01")

type Sender struct {
	address string
	iaddr   *net.TCPAddr
}

func NewSender(address string) *Sender {
	s := &Sender{address: address}
	if err := s.resolv(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Can't resolv zabbix address: %s\n", err.Error()))
		os.Exit(1)
	}
	return s
}

func (s *Sender) resolv() error {
	if iaddr, err := net.ResolveTCPAddr("tcp", s.address); err != nil {
		return err
	} else {
		s.iaddr = iaddr
	}
	return nil
}

func (s *Sender) connect() (*net.TCPConn, error) {
	if conn, err := net.DialTCP("tcp", nil, s.iaddr); err != nil { // TODO: timeout
		return nil, err
	} else {
		return conn, nil
	}
}

func (s *Sender) read(conn *net.TCPConn) ([]byte, error) {
	return ioutil.ReadAll(conn) // TODO: timeout
}

func (s *Sender) Send(packet *Packet) error {

	conn, err := s.connect()
	if err != nil {
		return err
	}

	buffer := append(zbxHeader, packet.DataLen()...)
	buffer = append(buffer, packet.Json()...)

	_, err = conn.Write(buffer)
	if err != nil {
		return err
	}

	_, err = s.read(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read zabbix response error: %s\n", err.Error())
	}

	return nil

}
