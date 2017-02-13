package messageHandler

import (
	"net"
	"time"
	"math/rand"
	"fmt"
	//"log"
	"counter"
	"serverConcPoolDyn/type"
	"log"
)

func MessageHandler(c net.Conn, in <-chan poolT.PoolT, abort chan<- int, id int, idThread int) {

	var v poolT.PoolT

	for {
		v = <-in
		t := rand.Int63n(10 * int64(time.Second))
		//fmt.Println("<< Sleep: ", t, "Id: ", id, " Sequence", pos)
		fmt.Println("<< Id:", id, ".", idThread, " Sequence:", v.Pos)

		time.Sleep(time.Duration(t))

		counter.Add()
		fmt.Println(">> Id:", id, ".", idThread, " Sequence:", v.Pos, " Data:", string(v.Data))

		data := v.Data
		_, err := c.Write(data)
		if err != nil {
			fmt.Println("Handler Write: ", err)
			abort <- 0
			//log.Fatal("Handler Write: ", err)
			log.Panic("Handler Write: ", err)
			return
		}

	}
}
