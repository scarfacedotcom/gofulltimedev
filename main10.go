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

func (s *Server) quit() {
	s.quitch <- struct{}{}
}

func (s *Server) loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			fmt.Println("quitting server")
			break mainloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		}
	}
	fmt.Println("Server shutting down....")
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received a message:", msg)
}

func main10() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 3)
		server.quit()
	}()

	server.start()

}
