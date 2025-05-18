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

type Option func(*Server)

func New(options ...Option) *Server {
	s := &Server{}
	for _, option := range options {
		option(s)
	}
	return s
}

func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) Option {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func (s *Server) Start() error {
	fmt.Println("server start :", s.Host, s.Port, s.timeout, s.maxConn)
	return nil
}
