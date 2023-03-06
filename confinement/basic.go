/*
Confinement is a technique which ensures there is information available from only
one concurrent process

In the following example, two functions are created.
  - chanOwner function creates and returns a read-only channel
  - consumer function, which takes a read-only channel as input

It can be noted in the following example that the data is provided by only the
consumer function. No other part of code can modify the data source but they can
ready from it.

	|‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|
	|		chanOwner		 |------------------|
	|________________________|					|
												|
												|
												|
												|
												↓
										|‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|
										|	  Channel      |←-----------|
										|__________________|			|
												|						|
												|						|
												|					 |‾‾‾‾‾|
												|					 |	X  | Consumer can not write to channel
	|‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|					|					 |_____|
	|		Consumer		 |←-----------------|						|
	|________________________|											|
				|														|
				|														|
				|_______________________________________________________|


*/

package confinement

import "fmt"

func Demo() {

	chanOwner := func() <-chan int {
		results := make(chan int, 5)

		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()

		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received %d\n", result)
		}
		fmt.Println("Done printing")
	}

	results := chanOwner()
	consumer(results)

}
