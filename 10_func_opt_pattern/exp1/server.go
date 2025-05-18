package main

import "fmt"

type Server struct {
	Host string
	Port int
}

func New(host string, port int) *Server {
	return &Server{host, port}
}
func (s *Server) ServerStart() error {
	fmt.Println("server start")
	return nil
}
