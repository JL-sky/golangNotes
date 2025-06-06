package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

func main() {
	person := Person{
		UserId:   "120",
		Username: "jack",
		Age:      18,
		Address:  "usa",
	}

	json, _ := json.Marshal(person)
	reader := bytes.NewReader(json)

	resp, err := http.Post("https://golang.org", "application/json;charset=utf-8", reader)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
