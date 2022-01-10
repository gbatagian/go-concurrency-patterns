// A package with functions about prime numbers generation,
// used to demonstrate concurency patterns examples.

package conc

import "math"

func IsPrime(integer int) bool {

	// trivial case for 2.
	if integer == 2 {
		return true
	}
	// trivial non prime numbers.
	if integer%2 == 0 || integer == 1 {
		return false
	}

	// check if prime.
	loop_range := int(math.Floor(math.Sqrt(float64(integer))))
	for i := 3; i <= loop_range; i += 2 {
		if integer%i == 0 {
			return false
		}
	}

	return true

}

func PrimesBetween(start int, max_number int) []int {
	// Returns a slice with all the primes between :start and :max_number.
	var primes []int

	if max_number <= start {
		return primes
	}

	loop_range := make([]int, max_number-start+1)
	for idx := range loop_range {

		idx += start
		if IsPrime(idx) {
			primes = append(primes, idx)
		}

	}

	return primes

}
