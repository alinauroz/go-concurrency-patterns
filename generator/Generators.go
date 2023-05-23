package generator

import (
	"fmt"
	"time"
)

func Repeat(stop <-chan interface{}, data []int) <-chan int {
	stream := make(chan int)

	go func() {
		defer close(stream)
		for {
			for _, e := range data {
				select {
				case <-stop:
					return
				case stream <- e:
				}
			}
		}
	}()

	return stream
}

func RepeatDemo() {
	data := []int{1, 2, 3, 4, 5}
	stop := make(chan interface{})
	stream := Repeat(stop, data)

	go func() {
		time.Sleep(2 * time.Second)
		close(stop)
	}()

	for i := range stream {
		fmt.Println(i)
	}

	fmt.Println("Generator stopped!")

}
