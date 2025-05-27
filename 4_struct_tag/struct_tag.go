package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Movie struct {
	Title  string   `json:title`
	Year   int      `json:tear`
	Price  int      `json:price`
	Actors []string `json:actors`
}

func ConvertJson() {
	// 创建 Movie 实例
	movie := Movie{
		Title:  "喜剧之王",
		Year:   2020,
		Price:  100,
		Actors: []string{"周星驰", "莫文蔚"},
	}

	// 1. 将 movie 转换为 JSON 并保存到文件
	jsonData, err := json.MarshalIndent(movie, "", "  ")
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}

	// 写入文件
	err = os.WriteFile("movie.json", jsonData, 0644)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
	fmt.Println("成功保存到 movie.json")

	// 2. 从文件读取 JSON 并解析，fileData是一个 []byte 类型的切片
	fileData, err := os.ReadFile("movie.json")
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	var myMovie Movie
	err = json.Unmarshal(fileData, &myMovie)
	if err != nil {
		fmt.Println("JSON 反序列化失败:", err)
		return
	}

	// 打印解析结果
	fmt.Printf("从文件读取的 Movie:\n%+v\n", myMovie)
}

type Config struct {
	Mysql     Mysql     `yaml:"mysql"`
	Redis     Redis     `yaml:"redis"`
	GeoliConf GeoliConf `yaml:"GeoliConf"`
	Users     []User    `yaml:"User"`
}

type Mysql struct {
	Url  string
	Port int
}

type Redis struct {
	Host string
	Port int
}

type GeoliConf struct {
	MinVersion []map[string]string `yaml:"min_version"`
}

type User struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
	Sex  string `yaml:"sex"`
}

func ParseYaml() {
	dataBytes, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("yaml 文件内容：\n", string(dataBytes))

	var config Config
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		fmt.Println("解析 YAML 失败：", err)
		return
	}

	// 输出解析结果
	fmt.Println("=== 解析后的结构体 ===")
	fmt.Printf("Mysql:\n\tUrl=%s\n\tPort=%d\n", config.Mysql.Url, config.Mysql.Port)
	fmt.Printf("Redis:\n\tHost=%s\n\tPort=%d\n", config.Redis.Host, config.Redis.Port)
	fmt.Println("GeoliConf.min_version:")
	for i, version := range config.GeoliConf.MinVersion {
		for k, v := range version {
			fmt.Printf("  第 %d 项：Key=%s，Value=%s\n", i+1, k, v)
		}
	}
	fmt.Println("Users:")
	for i, user := range config.Users {
		fmt.Printf("  第 %d 项：Name=%s，Age=%d，Sex=%s\n", i+1, user.Name, user.Age, user.Sex)
	}
}
func main() {
	ConvertJson()
	// ParseYaml()
}
