package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type WorkerDTO struct {
	depth  int
	parent int
}

func main() {
	timeStart := time.Now()
	var c = make(chan WorkerDTO, 10)
	wg := sync.WaitGroup{}
	defer func() {
		wg.Wait()
		defer close(c)
		timeElapsed := time.Since(timeStart)
		println("Elapsed time in ms:", timeElapsed.Milliseconds())
	}()

	depth := getDepth()
	numOfWorkers := 5

	for i := 0; i < numOfWorkers; i++ {
		go func(workId int) {
			for dto := range c {
				if dto.depth == 0 {
					wg.Done()
					continue
				}

				for i := 0; i < 5; i++ {
					time.Sleep(5 * time.Millisecond)

					fmt.Printf("Worker %d, parent %d, depth %d\n", workId, dto.parent, depth-dto.depth+1)
					if dto.depth > 0 {
						wg.Add(1)
						c <- WorkerDTO{depth: dto.depth - 1, parent: i}
					}
				}
				wg.Done()
			}
		}(i)
	}

	wg.Add(1)
	c <- WorkerDTO{depth: depth, parent: 0}
}

func getDepth() int {
	depth := 1
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		depth = i
	}
	return depth
}
