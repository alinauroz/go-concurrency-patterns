package generator

func toInt(done <-chan interface{}, stream <-chan interface{}) <-chan int {
	outputStream := make(chan int)

	go func() {
		defer close(outputStream)
		for v := range stream {
			select {
			case <-done:
				return
			case outputStream <- v.(int):
			}
		}
	}()

	return outputStream
}
