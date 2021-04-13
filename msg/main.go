package msg

import (
	"fmt"
	"time"
)

type MsgClient struct {
	ch        messageChannel
	cch       controlChannel
	connected bool
	buffer    []message
}

const (
	META_CONNECT = iota
	META_DISCONNECT
)

type message struct {
	content string
}

type control struct {
	meta int
}
type messageChannel chan message
type controlChannel chan control

func Initialize() MsgClient {
	c := MsgClient{}
	c.ch = make(messageChannel, 10)
	c.cch = make(controlChannel, 10)
	go c.handler()
	return c
}

func (c *MsgClient) handler() {
	fmt.Println("Handler started.")
	for {
		select {
		case msg := <-c.ch:
			c.handleMsg(msg)
		case cmsg := <-c.cch:
			c.handleControl(cmsg)
		case <-time.After(1 * time.Second):
			c.flush() // likely not needed, but kept in case the transport can re-establish connection itself.
		}
	}
}

func (c *MsgClient) Connect() {
	msg := control{meta: META_CONNECT}
	c.cch <- msg
}
func (c *MsgClient) Disconnect() {
	msg := control{meta: META_CONNECT}
	c.cch <- msg
}

func (c *MsgClient) handleControl(cmsg control) {
	switch cmsg.meta {
	case META_CONNECT:
		c.connected = true
		fmt.Println("Connected")
		c.flush()
		fmt.Println("(flush done)")
	case META_DISCONNECT:
		c.connected = false
		fmt.Println("Disconnected")
	}
}

func (c *MsgClient) handleMsg(msg message) {
	if c.connected == true {
		fmt.Printf("Got message: %s\n", msg.content)
	} else {
		fmt.Printf("Message buffered: %s\n", msg.content)
		c.buffer = append(c.buffer, msg)
	}
}

func (c *MsgClient) Send(content string) {
	msg := message{
		content: content,
	}
	c.ch <- msg

}

func (c *MsgClient) flush() {
	if len(c.buffer) == 0 {
		return
	}
	if c.connected == false {
		fmt.Println("Client not connected, unable to flush")
		return
	}
	fmt.Println("(flushing)")
	for _, msg := range c.buffer {
		c.handleMsg(msg)
	}
	c.buffer = []message{}

}
