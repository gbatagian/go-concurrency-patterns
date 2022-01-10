package conc

import (
	"fmt"
)

func BufferedChannel() {
	// A trivial bufferred channel example.

	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	fmt.Println(<-c)
	fmt.Println(<-c)

}

func GetPrimesFromChannel() {
	// A single channel example.

	c := make(chan int)

	l := 100_000_000_000_000 // Start from a large number so that IsPrime takes a bit to execute.
	r := 2_000
	m := l + r

	go ChannelPrimesInRange(c, l, m)

	for prime := range c {
		fmt.Println(prime)
	}
}

func GetPrimesFromMultipleChannels() {
	// A multiple channels example.

	c1 := make(chan int)
	c2 := make(chan int)

	l := 100_000_000_000_000 // Start from a large number so that IsPrime takes a bit to execute.
	r := 1_000
	m := l + r

	go ChannelPrimesInRange(c1, l, m)
	go ChannelPrimesInRange(c2, m+1, m+1+2*r) // Twice the range compared to to first Goroutine

	cnt1 := 0
	cnt2 := 0
	for {
		select {
		case prime1, channel1_open := <-c1:

			if !channel1_open {
				c1 = nil
			} else {
				cnt1 += 1
				fmt.Println("Prime from channel 1:", prime1)
			}

		case prime2, channel2_open := <-c2:

			if !channel2_open {
				c2 = nil
			} else {
				cnt2 += 1
				fmt.Println("Prime from channel 2:", prime2)
			}
		}

		if c1 == nil && c2 == nil {
			fmt.Printf("\n--> Total primes: %d\n", cnt1+cnt2)
			fmt.Printf("--> From channel 1: %d\n", cnt1)
			fmt.Printf("--> From channel 2: %d\n", cnt2)
			break
		}

	}

}

func ChannelPrimesInRange(c chan int, start int, end int) {

	if end <= start {
		return
	}

	r := make([]int, end-start+1)
	for idx := range r {
		idx += start
		if IsPrime(idx) {
			c <- idx
		}
	}
	close(c)

}

func ChannelsMainExample() {

	GetPrimesFromMultipleChannels()

}
