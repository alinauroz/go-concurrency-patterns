package channel_types

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	done := make(chan interface{})

	go func() {
		defer close(done)
		switch len(channels) {
		case 2:
			{
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			}
		default:
			{
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], done)...):
				}
			}
		}
	}()

	return done
}

func OrDemo() {
	sig := func(after time.Duration) <-chan interface{} {
		done := make(chan interface{})

		go func() {
			defer close(done)
			time.Sleep(after)
		}()

		return done
	}

	now := time.Now()

	<-or(
		sig(time.Second),
		sig(time.Minute),
		sig(time.Minute*2),
		sig(time.Minute*2),
		sig(time.Hour),
	)

	fmt.Println("done after", time.Since(now))
}
