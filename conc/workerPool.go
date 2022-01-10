package conc

import (
	"fmt"
	"math"
	"time"
)

func worker(jobs <-chan int, results chan<- int) {

	for number := range jobs {
		if IsPrime(number) {
			results <- number
		} else {
			results <- -1
		}
	}

}

func ConcPrimesBetween(start int, max_number int, workers int) []int {
	// Returns a slice with all the primes between :start and :max_number.
	// The primes are calculated concurrently using the worker pool technique.

	var primes []int

	if max_number <= start {
		return primes
	}

	jobs := make(chan int, max_number-start+1)
	results := make(chan int, max_number-start+1)

	workers_range := make([]int, workers)
	for range workers_range {
		go worker(jobs, results)
	}

	loop_range := make([]int, max_number-start+1)
	for idx := range loop_range {
		jobs <- idx
	}
	close(jobs)

	for j := 1; j <= max_number-start+1; j++ {
		number := <-results
		if number != -1 {
			primes = append(primes, number)
		}
	}

	return primes

}

func WorkerPoolMainExample() {

	low := 0
	max := 50_000_000
	workers := 10

	start := time.Now()
	primes := len(ConcPrimesBetween(low, max, workers))
	fmt.Println("\nMethod: Worker Pool")
	fmt.Printf("Primes in range [%d, %d]: %v\n", low, max, primes)
	t1 := time.Since(start)
	fmt.Printf("Time taken: %s\n", t1)

	start = time.Now()
	primes = len(PrimesBetween(low, max))
	fmt.Println("\nMethod: Standard (synchronous)")
	fmt.Printf("Primes in range [%d, %d]: %v\n", low, max, primes)
	t2 := time.Since(start)
	fmt.Printf("Time taken: %s\n", t2)

	fmt.Printf("\n--> %f%% Performane Difference\n", 100*math.Abs(t2.Seconds()-t1.Seconds())/math.Min(t1.Seconds(), t2.Seconds()))

}
