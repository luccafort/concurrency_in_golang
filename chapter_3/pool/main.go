package main

import (
	"fmt"
	"sync"
)

// p.60
func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}

//p.61
// func main() {
// 	var numCalcsCreated int
// 	calcPool := &sync.Pool{
// 		New: func() interface{} {
// 			numCalcsCreated++
// 			mem := make([]byte, 1024)
// 			return &mem
// 		},
// 	}

// 	// プールに4kb確保する
// 	calcPool.Put(calcPool.New())
// 	calcPool.Put(calcPool.New())
// 	calcPool.Put(calcPool.New())
// 	calcPool.Put(calcPool.New())

// 	const numWorkers = 1024 * 1024
// 	var wg sync.WaitGroup
// 	wg.Add(numWorkers)

// 	for i := numWorkers; i > 0; i-- {
// 		go func() {
// 			defer wg.Done()

// 			mem := calcPool.Get().(*[]byte)
// 			defer calcPool.Put(mem)

// 			// なにか面白いことを行う
// 			// ただしメモリに対して素早い処理が行われること
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Printf("%d calculators were created.", numCalcsCreated)
// }
