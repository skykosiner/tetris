package tcp

import (
	"io"
	"log"
	"net"
	"sync"
)

type TCP struct {
	Listener    net.Listener
	Connections []net.Conn
	Mutex       sync.Mutex
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

		t.Mutex.Lock()
		t.Connections = append(t.Connections, conn)
		t.Mutex.Unlock()

		go t.handleConnection(conn)
	}
}

func (t *TCP) handleConnection(conn net.Conn) {
	defer func() {
		t.removeConnection(conn)
		conn.Close()
	}()

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

func (t *TCP) removeConnection(conn net.Conn) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	for i, connection := range t.Connections {
		if connection == conn {
			t.Connections = append(t.Connections[:i], t.Connections[i+1:]...)
			break
		}
	}
}
