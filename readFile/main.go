package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//defer
	content, err := ioutil.ReadFile("/home/fmoglia/examples.deskto")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", content)
}
