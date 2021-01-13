package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

// p62
// $ go test -benchtime=10s -bench=.
func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	// $ go test -benchtime=10s -bench=.
	// goos: darwin
	// goarch: amd64
	// BenchmarkNetworkRequest-16            10        1004227648 ns/op
	// PASS
	// ok      _/Users/konishi.tatsuro/github.com/luccafort/concurrency_in_golang/chapter_3/pool       11.357s

	// unusing sync.Pool
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Printf("cannot accept connection: %v", err)
				continue
			}

			connectToService()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()

	// $ go test -benchtime=10s -bench=.
	// goos: darwin
	// goarch: amd64
	// BenchmarkNetworkRequest-16          1000          15962675 ns/op
	// PASS
	// ok      _/Users/konishi.tatsuro/github.com/luccafort/concurrency_in_golang/chapter_3/pool       45.648s

	// using sync.Pool
	go func() {
		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()

	return &wg
}

func init() {
	daemoStarted := startNetworkDaemon()
	daemoStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
