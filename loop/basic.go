/*
	DEFAULT case in for select loop

	In this example, a channel is created which will be closed after 5seconds. In PART I,
	a goroutine sleeps for 5 seconds and then closes the channel.

	In PART II, there is a for select loop which will run infinitely until there is something
	to be read on channel. This "for select" is going to take 5 seconds (simulated time.Sleep)
	which means it will stop the execution of code for 5 seconds. This 5 seconds of CPU time will
	be wasted.

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
