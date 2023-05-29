package pipeline

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alinauroz/go-concurrency-patterns/generator"
)

func getPrimeStream(done <-chan interface{}, inputStream <-chan int) <-chan int {
	outputStream := make(chan int)

	go func() {
		defer close(outputStream)

		for v := range inputStream {
			select {
			case <-done:
				return
			case outputStream <- func() int {
				for i := v; i > 0; i-- {
					for j := v - 1; j > 1; j-- {
						if j == 2 {
						}
						if j == 2 && i%j != 0 {
							return i
						}
						if i == j {
							continue
						}
						if i%j == 0 {
							break
						}
					}
					continue
				}
				return 1
			}():
			}
		}

	}()

	return outputStream
}

func WithoutFanOutFanIn() {
	start := time.Now()
	done := make(chan interface{})
	getRand := func() interface{} {
		return rand.Intn(500000000)
	}
	randStream := generator.LimitedRepeat(done, generator.RepeatFunc(done, getRand), 10)
	randIntStream := generator.ToInt(done, randStream)
	primeStream := getPrimeStream(done, randIntStream)

	for v := range primeStream {
		fmt.Println(v)
	}

	close(done)
	fmt.Println("Done, it took", time.Since(start))
}
