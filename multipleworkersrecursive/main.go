package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

type WorkerDTO struct {
	depth  int64
	parent int64
}

func main() {
	timeStart := time.Now()
	depth := getDepth()
	leafs := int64(5)
	numOfWorkers := 5000
	var c = make(chan WorkerDTO, int64(math.Pow(float64(leafs), float64(depth))))
	wg := sync.WaitGroup{}
	defer func() {
		wg.Wait()
		close(c)
		timeElapsed := time.Since(timeStart)
		println("Elapsed time in ms:", timeElapsed.Milliseconds())
	}()

	wg.Add(1)
	c <- WorkerDTO{depth: depth, parent: 0}

	for i := 0; i < numOfWorkers; i++ {
		go func(workId int) {
			for dto := range c {
				if dto.depth == 0 {
					wg.Done()
					continue
				}

				for i := int64(0); i < leafs; i++ {
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
}

func getDepth() int64 {
	depth := int64(1)
	if len(os.Args) > 1 {
		i, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			panic(err)
		}
		depth = i
	}
	return depth
}
