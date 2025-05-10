package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:title`
	Year   int      `json:tear`
	Price  int      `json:price`
	Actors []string `json:actors`
}

func ConvertJson() {
	movie := Movie{
		Title:  "喜剧之王",
		Year:   2020,
		Price:  100,
		Actors: []string{"周星驰", "莫文蔚"},
	}
	// 将movie转换为json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	fmt.Printf("jsonStr=%s\n", jsonStr)

	myMoive := Movie{}
	err = json.Unmarshal(jsonStr, &myMoive)
	if err != nil {
		fmt.Println("json unmarshal failed, err:", err)
		return
	}
	fmt.Printf("myMoive:%v\n", myMoive)
}

func main() {
	ConvertJson()
}
