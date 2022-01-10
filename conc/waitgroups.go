package conc

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func WaitgroupsMainExample() {

	wg.Add(1)
	go func() {
		start := time.Now()
		l := 0
		m := 1_000_000
		primes := PrimesBetween(l, m)
		fmt.Printf("--> Done with Group 1. Primes in range [%d, %d]: %d | Time taken: %s\n", l, m, len(primes), time.Since(start))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		start := time.Now()
		l := 0
		m := 10_000_000
		primes := PrimesBetween(l, m)
		fmt.Printf("--> Done with Group 2. Primes in range [%d, %d]: %d | Time taken: %s\n", l, m, len(primes), time.Since(start))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		start := time.Now()
		l := 0
		m := 10_000
		primes := PrimesBetween(l, m)
		fmt.Printf("--> Done with Group 3. Primes in range [%d, %d]: %d | Time taken: %s\n", l, m, len(primes), time.Since(start))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		start := time.Now()
		l := 0
		m := 100_000
		primes := PrimesBetween(l, m)
		fmt.Printf("--> Done with Group 4. Primes in range [%d, %d]: %d | Time taken: %s\n", l, m, len(primes), time.Since(start))
		wg.Done()
	}()

	wg.Wait()
}
