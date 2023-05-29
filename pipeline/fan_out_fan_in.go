package pipeline

import (
	"fmt"

	"github.com/alinauroz/go-concurrency-patterns/generator"
)

func getPrimeStream(done <-chan interface{}, inputStream <-chan int) <-chan int {
	outputStream := make(chan int)

	go func() {
		defer close(outputStream)
		for {
			for v := range inputStream {
				select {
				case <-done:
					return
				case outputStream <- func() int {
					return v
				}():
				}
			}
		}
	}()

	return outputStream
}

func WithoutFanOutFanIn() {
	done := make(chan interface{})
	rand := func() interface{} {
		return 100000
	}
	randStream := generator.LimitedRepeat(done, generator.RepeatFunc(done, rand), 10)
	randIntStream := generator.ToInt(done, randStream)
	primeStream := getPrimeStream(done, randIntStream)

	for v := range primeStream {
		fmt.Println(v)
	}
}
