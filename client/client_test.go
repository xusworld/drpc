package client

import (
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:10086")
	if err != nil {

	}

	conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))

}
