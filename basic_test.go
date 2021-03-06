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

	var pGenerated Primes

	for c := 0; c < ncases; c++ {

		// range = [l, u]
		min := 30
		max := 300
		l := r.Intn(min)
		u := r.Intn(max-min) + min
		fmt.Printf("Range: [%v,%v] \n", l, u)

		pGenerated = basic(l, u)

		if len(pGenerated) == 0 {
			t.Errorf("Error!")
		}

		// pValidation: list of primes between l and u in primes.txt
		// isEqual: true if pGenerated and pValidation are equal
		pValidation, isValid := validate(pGenerated, l, u)
		if isValid {
			fmt.Println("Results: ")
			printSlice(pGenerated)
			fmt.Printf("\t -- Results are correct! \n")
		} else {
			fmt.Println("Results: ")
			printSlice(pGenerated)
			fmt.Println("Validation: ")
			printSlice(pValidation)
			fmt.Printf("\t -- Results are not correct! \n")
		}
	}
}
