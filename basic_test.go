package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// # of test cases = [1, 10]
	ncases := r.Intn(10) + 1
	fmt.Printf("Basic -- Number of test cases: %v \n", ncases)

	var primes Primes

	for c := 0; c < ncases; c++ {

		// range = [l, u]
		min := 30
		max := 300
		l := r.Intn(min)
		u := r.Intn(max-min) + min
		fmt.Printf("Range: [%v,%v] \n", l, u)

		// True or false
		// p := r.Intn(2)
		// v := r.Intn(2)

		primes = basic(l, u)

		if len(primes) == 0 {
			t.Errorf("Error!")
		}

		// if p == 1 {
		//     printSlice(primes)
		//	}

		var pValidation Primes
		pValidation, isValid := validate(primes, l, u)
		if isValid {
			fmt.Println("Results: ")
			printSlice(primes)

			fmt.Printf("\t -- Results are correct! \n")
		} else {
			fmt.Println("Results: ")
			printSlice(primes)
			fmt.Println("Validation: ")
			printSlice(pValidation)

			fmt.Printf("\t -- Results are not correct! \n")
		}
	}
}
