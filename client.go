package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

func (c *client) read() {
	log.Println("inside read:", c)
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
			log.Println("message forwarded:", msg, c)
		} else {
			log.Println("bailing out, socket read err:", err)
			break
		}

	}
	c.socket.Close()
}

func (c *client) write() {
	log.Println("inside write:", c)
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("message written", msg, c)
			break
		} else {
			log.Println("bailing out, socket write err:", err)
		}
	}
	c.socket.Close()
}
