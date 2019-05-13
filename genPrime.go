package main

import (
	"fmt"
	"runtime"
	"time"
)

// Primes is the list of prime numbers
type Primes []int

// Algorithm enum
const (
	algoBasic  int = 1
	algoEratos int = 2
	algoSSE    int = 3
)

// Flags are the input options provided by the user
type Flags struct {
	lowerBound int
	upperBound int
	algorithm  int
	isPrint    bool
	validate   bool
	parallel   bool
}

func printSlice(s Primes) {
	fmt.Printf("List of prime numbers: %v \n", s)
}

func main() {

	// Results are written to this slice
	var primes Primes

	var inFlags Flags
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

			var primesPerCore Primes

			offset := inFlags.upperBound / ncores
			pPerCore := make(chan Primes, ncores)

			// Launch goroutines on all logical cores
			for i := 0; i < ncores; i++ {

				// Calculate offset values for each goroutine
				lOffset := i*offset + 1
				rOffset := (i + 1) * offset
				fmt.Printf("Offset values [%v, %v] for core: %v \n", lOffset, rOffset, i)

				go basicPar(pPerCore, lOffset, rOffset)
				primesPerCore = <-pPerCore

				fmt.Printf("Primes per core: %v", i)
				printSlice(primesPerCore)

				for j := 0; j < len(primesPerCore); j++ {
					primes = append(primes, primesPerCore[j])
				}
			}
		}
	case algoEratos:
		primes = eratosthenes(inFlags.lowerBound, inFlags.upperBound)
	case algoSSE:
		primes = sse(inFlags.lowerBound, inFlags.upperBound)
	}

	// Print measured time
	fmt.Printf("Elapsed time: %v for range: [%v,%v] using algorithm: %v \n",
		time.Since(start), inFlags.lowerBound, inFlags.upperBound, inFlags.algorithm)

	fmt.Printf("\t Number of primes: %v \n", len(primes))
	fmt.Printf("\t Largest prime: %v \n", primes[len(primes)-1])

	if inFlags.isPrint {
		printSlice(primes)
	}

	if inFlags.validate {
		isValid := validate(primes, inFlags.lowerBound, inFlags.upperBound)
		if isValid {
			fmt.Println(" -- Results are correct!")
		} else {
			fmt.Println(" -- Results are not correct!")
		}
	}
}
