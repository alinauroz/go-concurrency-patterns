/*
	A generator is a function which produces sequences of values on a
	channel.
*/

package generator

import "fmt"

func Generator() <-chan int {
	stream := make(chan int)
	arr := []int{1, 2, 3, 4, 5}

	go func() {
		defer close(stream)
		for i := 0; i < len(arr); i++ {
			stream <- arr[i]
		}
	}()

	return stream
}

func Demo() {
	stream := Generator()
	for v := range stream {
		fmt.Println(v)
	}
}
