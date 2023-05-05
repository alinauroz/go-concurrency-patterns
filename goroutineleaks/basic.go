/*
	A Goroutine Leak

	Introduction:
	In this demo, we are going to create a goroutine which will never be terminated. Such
	condition is known as goroutine routine.

	Problem:
	Undoubtedly, goroutines are lightweight and do not require a significant amount of memory.
	However, they still consume resources, and hence, it is not desirable for them to run indefinitely.

	In the following example, a nil channel is passed to doWork. Receiving a value from a nil
	channel is blocked forever. Therefore the goroutine will remain in memory until the process is
	terminated. The lifespan of this process is short but in real world, processed can be long and
	this leaked goroutine can create other goroutines.

	Solution:
	Goroutines are created by other goroutines. So there is always a parent goroutine, the main
	goroutine. During their execution, goroutines communicate with other goroutines.

*/

package goroutineleaks

import (
	"fmt"
	"math/rand"
)

func Demo() {
	// This code will cause go-routine leak
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("Work Done")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()

		return completed
	}

	doWork(nil)
	fmt.Println("Done")

}

func WriteDemo() {

	randStream := func() <-chan int {
		stream := make(chan int)

		go func() {
			defer fmt.Println("Closed")
			defer close(stream)
			for {
				stream <- rand.Int()
			}
		}()

		return stream
	}()

	for i := 0; i < 3; i++ {
		defer fmt.Println("Random value is", <-randStream)
	}

	fmt.Println("Main finished")

}
