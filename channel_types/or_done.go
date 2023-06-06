package channel_types

import (
	"fmt"
	"time"

	"github.com/alinauroz/go-concurrency-patterns/generator"
)

func OrDone(done <-chan interface{}, channel <-chan interface{}) <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-channel:
				if ok == false {
					return
				}

				select {
				case stream <- v:
				case <-done:
				}
			}
		}
	}()

	return stream
}

func OrDoneDemo() {
	done := make(chan interface{})
	data := []int{1, 2, 3}
	inStream := generator.Repeat(done, data)

	resultStream := OrDone(done, inStream)

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Closing...")
		close(done)
	}()

	for i := range resultStream {
		fmt.Println(i)
	}
}
