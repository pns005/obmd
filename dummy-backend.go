package main

import (
	"fmt"
	"io"
	"net"
)

// A "dummy" IpmiDialer, that rather than actually speaking Ipmi,
// Connects to the Addr it is passed via tcp, sends the info to
// the destination address, and then returns that connection.
// This is useful for experimentation.
type DummyIpmiDialer struct {
}

func (d *DummyIpmiDialer) DialIpmi(info *IpmiInfo) (io.ReadWriteCloser, error) {
	conn, err := net.Dial("tcp", info.Addr)
	if err != nil {
		return nil, err
	}
	_, err = fmt.Fprintln(conn, info)
	if err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}