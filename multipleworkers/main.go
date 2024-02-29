package main

import (
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	timeStart := time.Now()
	var c = make(chan int, 1000)
	wg := sync.WaitGroup{}
	defer func() {
		wg.Wait()
		defer close(c)
		timeElapsed := time.Since(timeStart)
		println("Elapsed time in ms:", timeElapsed.Milliseconds())
	}()

	numOfWorkers := 50
	for i := 0; i < numOfWorkers; i++ {
		go func() {
			for range c {
				// println(i)
				time.Sleep(10 * time.Millisecond)
				wg.Done()
			}
		}()
	}

	for i := 0; i < getItems(); i++ {
		wg.Add(1)
		c <- i
	}

	// This is a comment
	println("Hello, World!")
}

func getItems() int {
	items := 10
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		items = i
	}
	return items
}
