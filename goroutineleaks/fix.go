package goroutineleaks

import (
	"fmt"
	"time"
)

func FixByChannel() {

	doWork := func(strings <-chan string, done <-chan interface{}) <-chan interface{} {
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
	completed := doWork(nil, done)

	go func() {
		defer close(done)
		time.Sleep(1 * time.Second)
	}()

	<-completed
	fmt.Println("DONE")

}
