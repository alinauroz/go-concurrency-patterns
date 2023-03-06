/*
	A basic for select loop
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
