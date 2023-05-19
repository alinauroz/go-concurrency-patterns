package pipeline

import "fmt"

// provides understanding of the concept of line
func Idea() {

	add := func(arr []int, value int) []int {
		result := make([]int, len(arr))

		for i, v := range arr {
			result[i] = v + value
		}

		return result
	}

	multiply := func(arr []int, value int) []int {
		result := make([]int, len(arr))

		for i, v := range arr {
			result[i] = v * value
		}

		return result
	}

	list := []int{1, 2, 3}
	fmt.Println(
		multiply(
			add(list, 1),
			2,
		),
	)
}

// uses generator for memory optimizations
func PipelineWithGenerator() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	add := func(done <-chan interface{}, stream <-chan int, number int) <-chan int {
		addedStream := make(chan int)

		go func() {
			defer close(addedStream)
			for i := range stream {
				select {
				case <-done:
					return
				case addedStream <- (i + number):
				}
			}
		}()

		return addedStream
	}

	multiply := func(done <-chan interface{}, stream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range stream {
				select {
				case <-done:
					return
				case multipliedStream <- (i * multiplier):
				}
			}
		}()
		return multipliedStream
	}

	arr := []int{1, 2, 3}
	done := make(chan interface{})
	inputStream := generator(done, arr...)
	resultStream := multiply(done, add(done, multiply(done, inputStream, 2), 1), 10)
	for i := range resultStream {
		fmt.Println(i)
	}
}

//keywords: generators, pipelines, stages
