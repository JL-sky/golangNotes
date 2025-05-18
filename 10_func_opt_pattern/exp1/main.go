package main

import "log"

func main() {
	Server := New("localhost", 8080)
	if err := Server.ServerStart(); err != nil {
		log.Fatal(err)
	}
}
