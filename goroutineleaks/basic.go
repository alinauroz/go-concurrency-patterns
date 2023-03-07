/*
	A Goroutine Leak

	Introduction:
	In this demo, we are going to create a goroutine which will never be terminated. Such
	condition is known as goroutine routine.

	Problem:
	Undoubtedly, goroutines are lightweight and do not require a significant amount of memory.
	However, they still consume resources, and hence, it is not desirable for them to run indefinitely.

	Solution:
	Goroutines are created by other goroutines. So there is always a parent goroutine, the main
	goroutine. During their execution, goroutines communicate with other goroutines.

*/

package goroutineleaks

import "fmt"

func Demo() {

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
