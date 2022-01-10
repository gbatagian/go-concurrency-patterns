package conc

import (
	"fmt"
	"sync"
	"time"
)

func WaitgroupsMainExample() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		t_start := time.Now()
		start := 0
		end := 1_000_000
		primes := PrimesBetween(start, end)
		fmt.Printf("--> Done with Group 1. Primes in range [%d, %d]: %d | Time taken: %s\n", start, end, len(primes), time.Since(t_start))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		t_start := time.Now()
		start := 0
		end := 10_000_000
		primes := PrimesBetween(start, end)
		fmt.Printf("--> Done with Group 2. Primes in range [%d, %d]: %d | Time taken: %s\n", start, end, len(primes), time.Since(t_start))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		t_start := time.Now()
		start := 0
		end := 10_000
		primes := PrimesBetween(start, end)
		fmt.Printf("--> Done with Group 3. Primes in range [%d, %d]: %d | Time taken: %s\n", start, end, len(primes), time.Since(t_start))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		t_start := time.Now()
		start := 0
		end := 100_000
		primes := PrimesBetween(start, end)
		fmt.Printf("--> Done with Group 4. Primes in range [%d, %d]: %d | Time taken: %s\n", start, end, len(primes), time.Since(t_start))
		wg.Done()
	}()

	wg.Wait()
}
