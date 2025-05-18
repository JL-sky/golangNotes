package main

import (
	"log"
	"time"
)

func main() {
	// Server := New("localhost", 8080)
	Server := NewWithTimeout("localhost", 8080, 10*time.Second)
	if err := Server.ServerStart(); err != nil {
		log.Fatal(err)
	}
}
