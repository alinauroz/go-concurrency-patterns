package pipeline

import (
	"fmt"
	"time"

	"github.com/alinauroz/go-concurrency-patterns/generator"
)

func MeasureTime(process string) func() {
	fmt.Printf("Start %s\n", process)
	start := time.Now()
	return func() {
		fmt.Printf("Time taken by %s is %v\n", process, time.Since(start))
	}
}

func QueueDemo() {

	fmt.Println("Part I: Queue Demo Without Buffer")

	done := make(chan interface{})
	inp := []int{0, 0, 0}

	sleep := func(done <-chan interface{}, duration int, inputStream <-chan interface{}, name string) <-chan interface{} {
		stream := make(chan interface{})

		go func() {
			defer close(stream)
			defer MeasureTime(name)()
			for {
				select {
				case <-done:
					return
				case v, ok := <-inputStream:
					if !ok {
						return
					}
					time.Sleep(time.Duration(duration) * time.Second)
					stream <- v
				}
			}
		}()

		return stream
	}

	zeros := generator.LimitedRepeat(done, generator.Repeat(done, inp), 3)
	short := sleep(done, 1, zeros, "Short Without Buffer")
	long := sleep(done, 3, short, "Long Without Buffer")

	for v := range long {
		fmt.Println("Received value from Long wihtout buffer: ", v)
	}

}
