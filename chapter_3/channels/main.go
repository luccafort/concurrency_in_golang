package main

import "fmt"

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

// invalid channels
// func main() {
// 	// $ go run main.go
// 	// # command-line-arguments
// 	// ./main.go:29:2: invalid operation: <-writeStream (receive from send-only type chan<- interface {})
// 	// ./main.go:30:13: invalid operation: readStream <- struct {} literal (send to receive-only type <-chan interface {})
// 	writeStream := make(chan<- interface{})
// 	readStream := make(<-chan interface{})

// 	<-writeStream
// 	readStream <- struct{}{}
// }

// p.68
// dead lock
// func main() {
// 	// $ go run main.go
// 	// fatal error: all goroutines are asleep - deadlock!
// 	//
// 	// goroutine 1 [chan receive]:
// 	// main.main()
// 	//         /Users/konishi.tatsuro/github.com/luccafort/concurrency_in_golang/chapter_3/channels/main.go:50 +0x7c
// 	// exit status 2
// 	stringStream := make(chan string)
// 	go func() {
// 		if 0 != 1 {
// 			return
// 		}
// 		stringStream <- "Hello, channels!"
// 	}()

// 	fmt.Println(<-stringStream)
// }
