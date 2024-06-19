package main

import "fmt"

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func newSever() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string),
	}
}

func (s *Server) start() {
	fmt.Println("starting server")
}

func (s *Server) loop() {
	for {
		select {
		case <-s.quitch:
		case msg := <-s.msgch:
			_ = msg
		}
	}
}

func main() {

}
