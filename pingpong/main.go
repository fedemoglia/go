package main

import (
	"fmt"
	"math/rand"
	"time"
)

// chan<- para evitar errores futuros en el código al tratar de leer el canal en este código
func ping(ball chan<- int, action chan<- string) {
	ball <- 1
	action <- "Player ping"
}

func pong(ball chan<- int, action chan<- string) {
	ball <- 2
	action <- "Player pong"
}

func referee(action <-chan string) {
	for {
		fmt.Println("Action: ", <-action)
	}
}

func pingpong() {
	ball := make(chan int)
	action := make(chan string)
	go referee(action)
	go ping(ball, action)
	for {
		value := <-ball
		switch value {
		case 1:
			go pong(ball, action)
		case 2:
			go ping(ball, action)
		}
	}
}

func main() {
	var key string
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	seconds := r1.Intn(10)
	fmt.Printf("Comienza el juego de ping pong por %d segundos\n", seconds)
	fmt.Println("Presione ENTER para comenzar")
	fmt.Scanf("%d", &key)
	go pingpong()
	time.Sleep(time.Duration(seconds) * time.Second)
}
