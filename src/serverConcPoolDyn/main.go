// Server that creates a pool of coroutine for handling each a received message.
// The set of pools is dynamic. Creates new coroutines if are few
// The handler of a connection is not properly detected the failed of a
// client

package main

import
(
	"log"
	"net"
	"serverConcPoolDyn/connectionHandler"
	"fmt"
)




func main() {

	i := 0

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
			fmt.Println("accept error:", err)
			log.Fatal("accept error:", err)
		}
		i++
		go connectionHandler.ConnHandler(fd, i)
	}
}
