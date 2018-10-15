package native

import (
	"fmt"
	"net"

	"github.com/cjburchell/yasls-client-go"
)

type connection struct {
	conn net.Conn
}

func newConnection(address string, port int) (*connection, error) {
	log.Printf("Connecting to %s:%d", address, port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return nil, err
	}

	return &connection{conn: conn}, nil
}

func (c connection) Write(data []byte) error {
	log.Printf("Write Buffer")
	_, err := c.conn.Write(data)
	return err
}
func (c connection) Read(size int) ([]byte, int, error) {
	log.Printf("Read Buffer")
	buf := make([]byte, size)
	read, err := c.conn.Read(buf)
	log.Printf("Read %d %v", read, buf)
	return buf, read, err
}
func (c connection) Disconnect() {
	c.conn.Close()
}
