/*
	Fixing Goroutine Leak Using a Channel

	In Demo function it can be seen how does a goroutine leak look. The same function
	is used in this example but it has been changed a bit. A done channel is passed to
	the doWork function.

	The program has been divided into 3 parts. In first part, doWork is defined and
	called. In Part II, there is a goroutine which closes the done channel after 1 second.
	So it does not matter how long the process will take to terminate, the goroutine will
	return after after second when the done channel is closed by the goroutine in PART II.

*/

package goroutineleaks

import (
	"fmt"
	"math/rand"
	"time"
)

func FixByChannel() {

	//PART I
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer close(completed)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()

		return completed
	}

	done := make(chan interface{})
	completed := doWork(done, nil)

	// PART II
	go func() {
		defer close(done)
		time.Sleep(1 * time.Second)
	}()

	// PART II
	<-completed
	fmt.Println("DONE")

}

func FixWriteLeakByChannel() {

	getRandStream := func(done <-chan interface{}) <-chan int {
		stream := make(chan int)

		go func() {
			defer fmt.Println("Closed")
			defer close(stream)
			for {
				select {
				case stream <- rand.Int():
				case <-done:
					return
				}

			}
		}()

		return stream
	}

	done := make(chan interface{})
	randStream := getRandStream(done)

	for i := 0; i < 3; i++ {
		fmt.Println("Random value is", <-randStream)
	}
	close(done)

	fmt.Println("Main finished")

}
