package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("starting server")
	s.loop()
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) loop() {
	for {
		select {
		case <-s.quitch:
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
		}
	}
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received a message:", msg)
}

func main() {
	server := newServer()
	go server.start()

	server.sendMessage("testing scar face")
	time.Sleep(time.Second * 2)

}
