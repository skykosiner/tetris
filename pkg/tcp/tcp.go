package tcp

import (
	"io"
	"log"
	"net"
)

type TCP struct {
	Listener    net.Listener
	Connections []net.Conn
}

func (t *TCP) Start() {
	listener, err := net.Listen("tcp", "localhost:42069")
	if err != nil {
		log.Fatal("Error starting TCP server: ", err)
	}

	t.Listener = listener
	defer t.Listener.Close()

	log.Println("Server is listening on port 42069")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err)
			continue
		}

		t.Connections = append(t.Connections, conn)
		go t.handleConnection(conn)
	}
}

func (t *TCP) handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by client")
			} else {
				log.Println("Error reading from connection: ", err)
			}
			break
		}

		log.Println("Received: ", string(buffer[:n]))
		t.ToAll("PENIS")
	}
}

func (t *TCP) ToAll(msg string) {
	for _, connection := range t.Connections {
		_, err := connection.Write([]byte(msg))
		if err != nil {
			log.Println("Error writing to connection: ", err)
		}
	}
}
