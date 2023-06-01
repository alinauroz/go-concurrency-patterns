package pipeline

import (
	"fmt"
	"math/rand"
	"sync"
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

func fan_in(done <-chan interface{}, channels []<-chan int) <-chan interface{} {
	multiplexed := make(chan interface{})
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, c := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				select {
				case <-done:
					return
				case multiplexed <- v:
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(multiplexed)
	}()

	return multiplexed
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

func WithFanOutFanIn() {
	start := time.Now()
	done := make(chan interface{})
	getRand := func() interface{} {
		return rand.Intn(500000000)
	}
	randStream := generator.LimitedRepeat(done, generator.RepeatFunc(done, getRand), 10)
	randIntStream := generator.ToInt(done, randStream)

	primeCalculators := make([]<-chan int, 8)
	for i := 0; i < 8; i++ {
		primeCalculators[i] = getPrimeStream(done, randIntStream)
	}

	primeStream := fan_in(done, primeCalculators)

	for v := range primeStream {
		fmt.Println(v)
	}

	close(done)
	fmt.Println("Done, it took", time.Since(start))
}
