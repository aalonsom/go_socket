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

type poolT struct{
	data  []byte
	pos   int
}


func handler(c net.Conn, in <-chan poolT, id int, idThread int) {

	var v poolT

	for {
		v = <-in
		t := rand.Int63n(10 * int64(time.Second))
		//fmt.Println("<< Sleep: ", t, "Id: ", id, " Sequence", pos)
		fmt.Println("<< Id:", id, ".", idThread, " Sequence:", v.pos)

		time.Sleep(time.Duration(t))

		counter.Add()
		fmt.Println(">> Id:", id, ".", idThread, " Sequence:", v.pos, " Data:", string(v.data))

		data := v.data
		_, err := c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}

	}
}

func echoServer(c net.Conn, id int) {
	nPool    := 20 // numnber of threads in the pool
	idThread := 0  // id of a thread in the pool
	pos      := 0; // id of a message got
	var v poolT

	channel := make(chan poolT)

	// Create the pool with the threads
	for i := 0; i < nPool; i++  {
		idThread = i // Just in case additional threads are added in the future
		go handler(c, channel, id, idThread)
	}

	//When a message is received, is forwarded to a handler
	for {
		// get socket
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		pos ++
		v.pos  = pos
		v.data = buf[0:nr]

		channel <- v
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
