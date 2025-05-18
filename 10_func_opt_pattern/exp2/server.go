package main

import (
	"fmt"
	"time"
)

type Server struct {
	Host    string
	Port    int
	timeout time.Duration
	maxConn int
}

func New(host string, port int) *Server {
	return &Server{host, port, time.Minute, 100}
}
func NewWithTimeout(host string, port int, timeout time.Duration) *Server {
	return &Server{host, port, timeout, 100}
}

func NewWithTimeoutAndMaxConn(host string, port int, timeout time.Duration, maxConn int) *Server {
	return &Server{host, port, timeout, maxConn}
}
func (s *Server) ServerStart() error {
	fmt.Println("server start : ", s.Host, s.Port, s.timeout, s.maxConn)
	return nil
}
