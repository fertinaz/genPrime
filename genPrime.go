package main

import (
	"fmt"
	"runtime"
	"time"
)

// Primes is the list of prime numbers
type Primes []int

var inFlags Flags

func main() {

	// Results are written to this slice
	var primes Primes

	parseFlags(&inFlags)

	var ncores int
	if inFlags.parallel {
		ncores = runtime.NumCPU()
		// fmt.Printf("Number of cores: %v \n", ncores)
	} else {
		ncores = 1
	}

	// Start measuring
	start := time.Now()

	// Run prime number detection
	switch inFlags.algorithm {
	case algoBasic:
		if ncores == 1 {
			primes = basic(inFlags.lowerBound, inFlags.upperBound)
		} else {

			// Some arbitrary number for # of jobs
			// invoked from the worker pool
			n := 10
			jobs := make(chan int, n)
			results := make(chan Primes, n)

			delta := inFlags.upperBound / n
			// delta := (inFlags.upperBound - inFlags.lowerBound) / n
			// for c := 0; c < ncores; c++ {
			go worker(jobs, results, delta)
			//}

			for i := 0; i < n; i++ {
				jobs <- i
			}

			close(jobs)

			for j := 0; j < n; j++ {
				printSlice(<-results)
			}
		}
	case algoEratos:
		primes = eratosthenes(inFlags.lowerBound,
			inFlags.upperBound)
	case algoSSE:
		primes = sse(inFlags.lowerBound, inFlags.upperBound)
	}

	// Print measured time
	fmt.Printf("Elapsed time: %v for range: [%v,%v] using algorithm: %v \n",
		time.Since(start),
		inFlags.lowerBound,
		inFlags.upperBound,
		inFlags.algorithm)

	// Print brief information about results
	fmt.Printf("\t Number of primes: %v \n", len(primes))
	fmt.Printf("\t Largest prime: %v \n", primes[len(primes)-1])

	// Print primes found
	if inFlags.isPrint {
		fmt.Println("Results: ")
		printSlice(primes)
	}

	// Check results
	if inFlags.validate {
		largestPrimeInDB := 1000033
		if inFlags.upperBound > largestPrimeInDB {
			fmt.Println(" -- Cannot validate values larger than 1 million!")
		} else {
			var pValidation Primes
			pValidation, isValid := validate(primes,
				inFlags.lowerBound,
				inFlags.upperBound)

			fmt.Println("Validation: ")
			printSlice(pValidation)

			if isValid {
				fmt.Println("-- Results are correct!")
			} else {
				fmt.Println("-- Results are not correct!")
			}
		}
	}
}
