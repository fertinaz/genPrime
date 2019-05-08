package main

import "math"

/*
	Segmented version of sieve of eratosthenes
	Input:
		l: lower bound of the input range -- int
		u: upper bound of the input range -- int
	Return:
		p: List of primes between input range
*/
func sse(l, u int) []int {

	var p []int
	var pFinal []int

	// Find primes from 2 up to sqrt(upperBound)
	delta := int(math.Sqrt(float64(u)))
	p = eratosthenes(2, delta)

	for i := 0; i < len(p); i++ {
		if p[i] >= l {
			pFinal = append(pFinal, p[i])
		}
	}

	// Divide into segments
	low := delta
	high := 2 * delta

	for low <= u {

		if high >= u {
			high = u
		}

		// Mark primes in current segment
		mark := make([]bool, delta+1)
		for i := 0; i < delta+1; i++ {
			mark[i] = true
		}

		// Use initial prime list to find primes in this segment
		for i := 0; i < len(p); i++ {

			// Find the minimum number in current segment which is a multiple of p[i]
			lowLim := int(math.Floor(float64(low/p[i])) * float64(p[i]))

			if lowLim < low {
				lowLim += p[i]
			}

			// Mark multiples of prime[i] in this segment
			for j := lowLim; j <= high; j += p[i] {
				mark[int(j)-low] = false
			}
		}

		// Numbers which are not marked as false are prime
		for i := low; i <= high; i++ {
			if mark[i-low] && i >= l {
				pFinal = append(pFinal, i)
			}
		}

		low += delta
		high += delta

	} // End outer for

	return pFinal
}
