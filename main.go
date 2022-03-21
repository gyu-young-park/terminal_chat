package main

import (
	"log"
	"net"

	"github.com/gyu-young-park/terminal_chat/chat"
)

func main() {
	s := chat.NewServer()
	go s.Run()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("unable to start server: " + err.Error())
	}

	defer listener.Close()
	log.Printf("started server on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		c := s.NewClient(conn)
		go c.ReadInput()
	}
}
