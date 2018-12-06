package native

import (
	"fmt"
	"net"
)

type connection struct {
	conn net.Conn
}

func newConnection(address string, port int) (*connection, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return nil, err
	}

	return &connection{conn: conn}, nil
}

func (c connection) Write(data []byte) error {
	_, err := c.conn.Write(data)
	return err
}
func (c connection) Read(size int) ([]byte, int, error) {
	buf := make([]byte, size)
	read, err := c.conn.Read(buf)
	return buf, read, err
}
func (c connection) Disconnect() {
	c.conn.Close()
}
