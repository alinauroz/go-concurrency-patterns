package pipeline

func getPrimeStream(done <-chan interface{}, inputStream <-chan int) <-chan int {
	outputStream := make(chan int)

	go func() {
		defer close(outputStream)
		for {
			for v := range inputStream {
				select {
				case <-done:
					return
				case outputStream <- func() int {
					return v
				}():
				}
			}
		}
	}()

	return outputStream
}

func WithoutFanOutFanIn() {

}
