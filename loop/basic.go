/*
	DEFAULT case in for select loop

	In this example, a channel is created which will be closed after 5seconds. In PART I,
	a goroutine sleeps for 5 seconds and then closes the channel.

	PART II contains a "for select" loop that will continuously run until there is data available
	on the channel. However, this "for select" loop will cause a delay of 5 seconds due to the
	simulated time.Sleep, effectively halting the code's execution and wasting 5 seconds of CPU time.

	This problem can be solved by default case. When none of the conditions for the select statement
	is fulfilled, default case is executed.
*/

package loop

import (
	"fmt"
	"time"
)

func Demo() {

	c := make(chan int)

	//PART I
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	//PART II
loop:
	for {
		select {
		case <-c:
			fmt.Println("Done")
			break loop
		default:
			fmt.Println("Doing other work")
			time.Sleep(1 * time.Second)
		}
	}
}
