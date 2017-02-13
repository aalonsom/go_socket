// Server that creates a pool of threads for handling each a received message.
// The set of pools is fixed.
// The handler of a connection is not properly detected the failed of a
// client


package main

import
(
	"log"
	"net"
	"testCountWait/counter"
	"fmt"
	"math/rand"
	"time"
)

func handler(c net.Conn, data []byte, id int, pos int) {

	t := rand.Int63n(10 * int64(time.Second))
	//fmt.Println("<< Sleep: ", t, "Id: ", id, " Sequence", pos)
	fmt.Println("<< Id: ", id, " Sequence", pos)

	time.Sleep(time.Duration(t))

	counter.Add()
	fmt.Println(">> Id: ", id, " Sequence: ", pos, " Data: ", string(data))

	_, err := c.Write(data)
	if err != nil {
		log.Fatal("Write: ", err)
	}

}

func echoServer(c net.Conn, id int) {
	pos := 0;
	for {
		// get socket
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		pos ++
		go handler(c, data, id, pos)

		// process socket
	}
	fmt.Println("Connecition ", id, " messages in server: ", counter.Get())
}

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
			log.Fatal("accept error:", err)
		}
		i++
		go echoServer(fd, i)
	}
}
