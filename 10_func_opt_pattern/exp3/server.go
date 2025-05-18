package main

import (
	"fmt"
	"time"
)

type Server struct {
	cfg Config
}

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
	MaxConn int
}

func New(cfg Config) *Server {
	return &Server{cfg}
}

func (s *Server) ServerStart() error {
	fmt.Println("server start :", s.cfg.Host, s.cfg.Port, s.cfg.Timeout, s.cfg.MaxConn)
	return nil
}
