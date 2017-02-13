package main

import (
	"log"
	"net"
)

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("Server got:", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func main() {

	// The server is going to receive from a tcp address and a port
	//l, err := net.Listen("unix", "/tmp/echo.sock")
	l, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		// Waits for a client and create the socket connection
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(fd)
	}
}
