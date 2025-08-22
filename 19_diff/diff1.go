package main

import (
	"encoding/json"
	"fmt"

	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

func Compare1(json1, json2 string) {
	differ := gojsondiff.New()
	d, err := differ.Compare([]byte(json1), []byte(json2))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if d.Modified() {
		fmt.Println("JSONs are different")
		var aJson map[string]interface{}
		json.Unmarshal([]byte(json1), &aJson)

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}
		formatter := formatter.NewAsciiFormatter(aJson, config)
		diffString, err := formatter.Format(d)
		if err != nil {
			fmt.Println("Error formatting diff:", err)
			return
		}
		fmt.Println(diffString)
	} else {
		fmt.Println("JSONs are identical")
	}
}
