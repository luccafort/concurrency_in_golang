package main

import (
	"fmt"
)

// p.66
// func main() {
// 	var receivedChan <-chan interface{}
// 	var sendChan chan<- interface{}
// 	dataStream := make(chan interface{})

// 	// 正しい記述
// 	receivedChan = dataStream
// 	sendChan = dataStream
// }

// p.67
func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello, channels!"
	}()

	fmt.Println(<-stringStream)
}
