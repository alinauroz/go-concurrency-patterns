package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func Repeat(stop <-chan interface{}, data []int) <-chan interface{} {
	stream := make(chan interface{})

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

func LimitedRepeat(done <-chan interface{}, stream <-chan interface{}, limit int) <-chan interface{} {
	limitedStream := make(chan interface{})

	go func() {
		defer close(limitedStream)
		for i := 0; i < limit; i++ {
			select {
			case <-done:
				return
			case limitedStream <- <-stream:
			}
		}
	}()

	return limitedStream
}

func LimitedRepeatDemo() {
	data := []int{1, 2, 3, 4, 5}
	stop := make(chan interface{})
	stream := Repeat(stop, data)
	limitedStream := LimitedRepeat(stop, stream, 10)

	go func() {
		time.Sleep(2 * time.Second)
		close(stop)
	}()

	for i := range limitedStream {
		fmt.Println(i)
	}

	fmt.Println("Generator stopped!")
}

func RepeatFunc(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func RepeatFuncDemo() {
	done := make(chan interface{})

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	stream := RepeatFunc(done, func() interface{} {
		return rand.Int()
	})

	for i := range stream {
		fmt.Println(i)
	}

}
