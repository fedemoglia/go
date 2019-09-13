package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/properties"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var client = &http.Client{Timeout: 10 * time.Second}

	url := getURL()

	resp, err := client.Get(url)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(resp.Body)

	var post []Post

	err = json.NewDecoder(resp.Body).Decode(&post)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(post[0])
}

func getURL() string {
	var prop = properties.MustLoadFile("properties.properties", properties.UTF8)
	return prop.GetString("url", "http://google.com")
}
