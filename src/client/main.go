// Server that creates a coroutine for handling each a received message.
// The handler of a connection is not properly detected the failed of a
// client. This occurs from writing a message to a client that has failed

package main

import (
"io"
"log"
"net"
"time"
	"fmt"
	"strconv"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("               Client got:", string(buf[0:n]))
	}
}

func main() {
	pos := 0
	//c, err := net.Dial("unix", "/tmp/echo.sock")
	c, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	go reader(c)
	for {	pos ++
		fmt.Println("Sequence: ", strconv.Itoa(pos))
		//_, err := c.Write([]byte("at"))
		_, err := c.Write([]byte(strconv.Itoa(pos)))
		if err != nil {
			log.Fatal("write error:", err)
			break
		}
		time.Sleep(1e9)
	}
}