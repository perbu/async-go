package main

import (
	"fmt"
	"github.com/perbu/async-go/msg"
	"time"
)

func main() {
	fmt.Println("Starting up")

	conn := msg.Initialize()

	conn.Send("Message 1")
	conn.Send("Message 2")
	time.Sleep(1000 * time.Millisecond)
	conn.Connect()
	time.Sleep(10 * time.Millisecond)
	conn.Send("Message 3")
	time.Sleep(2 * time.Second)

}
