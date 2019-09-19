package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var done = make(chan struct{})

func producer(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	for i := 1; i <= 10; i ++ {
		select {
		case <-done:   // close goroutine when close done channel
			return
		default:
			time.Sleep(3*time.Second)
			ch<- i
		}
	}
}

func consumer(wg *sync.WaitGroup, id int, ch <-chan int) {
	defer wg.Done()
loop:
	for {
		select {
		case v := <-ch:
			time.Sleep(1*time.Second)
			fmt.Printf("%d: %d\n", id, v)
		case <-time.Tick(1*time.Minute):      // 计时器
			fmt.Println("wait for data...")
			break loop
		case <-done:    // close goroutine when close done channel
			fmt.Printf("terminated single!")
			break loop
		}
	}
	fmt.Printf("exit %d\n", id)
}

func main()  {
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	// terminated single
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go consumer(&wg, i, ch)
	}
	wg.Add(1)
	go producer(&wg, ch)
	wg.Wait()
	close(ch)
	fmt.Println("Done!")
}
