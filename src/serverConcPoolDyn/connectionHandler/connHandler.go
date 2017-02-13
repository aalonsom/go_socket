package connectionHandler

import (
	"net"
	"fmt"
	//"counter"
	"serverConcPoolDyn/messageHandler"
	"serverConcPoolDyn/type"
	//"log"
)

func ConnHandler(c net.Conn, id int) {
	nPool    := 3 // numnber of threads in the pool
	idThread := 0  // id of a thread in the pool
	pos      := 0; // id of a message got
	var v poolT.PoolT

	channel := make(chan poolT.PoolT)
	abort   := make(chan int)

	// Create the pool with the threads
	for i := 0; i < nPool; i++  {
		idThread = i
		go messageHandler.MessageHandler(c, channel, abort, id, idThread)
	}

	//When a message is received, is forwarded to a handler
	for {
		// get socket
		buf := make([]byte, 512)
		// As it is blocked waiting for message, it is the first coroutine
		// that detects that the connection is closed
		nr, err := c.Read(buf)
		if err != nil {
			fmt.Println("Connection Write: ", err)
			return
		}
		pos ++
		v.Pos = pos
		v.Data = buf[0:nr]

		// If all coroutines are being used, new ones are created
		select {
			//Because of the previous comment, this isn ot called.
			case <-abort:
				fmt.Println("ID:", id, "finishing")
				return
			case channel <- v:
			default:
				idThread++
				go messageHandler.MessageHandler(c, channel, abort, id, idThread)
				channel <- v
		}
	}
}
