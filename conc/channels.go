package conc

import (
	"fmt"
)

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

func GetPrimesFromChannel() {
	// A single channel example.

	channel := make(chan int)

	start := 100_000_000_000_000 // Start from a large number so that IsPrime takes a bit to execute.
	rng := 2_000
	end := start + rng

	go ChannelPrimesInRange(channel, start, end)

	for prime := range channel {
		fmt.Println(prime)
	}
}

func GetPrimesFromMultipleChannels() {
	// A multiple channels example.

	channel1 := make(chan int)
	channel2 := make(chan int)

	start := 100_000_000_000_000 // Start from a large number so that IsPrime takes a bit to execute.
	rng := 1_000
	end := start + rng

	go ChannelPrimesInRange(channel1, start, end)
	go ChannelPrimesInRange(channel2, end+1, end+1+2*rng) // Twice the range compared to to first Goroutine

	cnt1 := 0
	cnt2 := 0
	for {
		select {
		case prime1, channel1_open := <-channel1:

			if !channel1_open {
				channel1 = nil
			} else {
				cnt1 += 1
				fmt.Println("Prime from channel 1:", prime1)
			}

		case prime2, channel2_open := <-channel2:

			if !channel2_open {
				channel2 = nil
			} else {
				cnt2 += 1
				fmt.Println("Prime from channel 2:", prime2)
			}
		}

		if channel1 == nil && channel2 == nil {
			fmt.Printf("\n--> Total primes: %d\n", cnt1+cnt2)
			fmt.Printf("--> From channel 1: %d\n", cnt1)
			fmt.Printf("--> From channel 2: %d\n", cnt2)
			break
		}

	}

}

func ChannelsMainExample() {
	GetPrimesFromMultipleChannels()
}

func BufferedChannel() {
	// A trivial bufferred channel example.

	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	fmt.Println(<-c)
	fmt.Println(<-c)

}
