package main

import (
	"log"
	"time"
)

func main() {
	config := Config{
		Host:    "localhost",
		Port:    8080,
		Timeout: 10 * time.Second,
	}
	Server := New(config)

	if err := Server.ServerStart(); err != nil {
		log.Fatal(err)
	}
}
