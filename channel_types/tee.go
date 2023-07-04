package channel_types

import (
	"fmt"
	"sync"

	"github.com/alinauroz/go-concurrency-patterns/generator"
)

func Tee(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)
		for val := range OrDone(done, in) {
			var out1, out2 = out1, out2 //variable shadowing
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil // making it nil, so on next iteration, value is not passed to this channel
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()

	return out1, out2
}

func TeeDemo() {
	inArray := []int{1, 2, 3}
	done := make(chan interface{})
	in := generator.LimitedRepeat(done, generator.Repeat(done, inArray), 10)

	wg := sync.WaitGroup{}
	wg.Add(2)

	out1, out2 := Tee(done, in)

	go func() {
		defer wg.Done()
		for v1 := range out1 {
			fmt.Println("V1: ", v1)
		}
	}()

	go func() {
		defer wg.Done()
		for v2 := range out2 {
			fmt.Println("V2: ", v2)
		}
	}()

	wg.Wait()
}
