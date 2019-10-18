package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Flags are the input options provided by the user
type Flags struct {
	lowerBound int
	upperBound int
	algorithm  int
	isPrint    bool
	measure    bool
	validate   bool
	parallel   bool
}

// Global algorithm enum
const (
	algoBasic  int = 1
	algoEratos int = 2
	algoSSE    int = 3
)

//	Function splits input string for range. Exits if the input format is wrong.
//	Input:
//		Range entered by the user as string
//	Return:
//		lower and upper bound values as integers
func getRange(str string) (l, u int) {

	const MaxInt uint64 = 1<<64 - 1

	s := strings.Split(str, ",")

	l, err := strconv.Atoi(s[0])
	if err == nil {
		if l < 1 || uint64(l) > MaxInt {
			fmt.Printf("Lower value must be a non-negative integer and should not exceed: %v \n", MaxInt)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Format error: %v \n", err)
		os.Exit(1)
	}

	u, err = strconv.Atoi(s[1])
	if err == nil {
		if u < 1 || u < l || uint64(u) > MaxInt {
			fmt.Printf("Upper value must be a non-negative integer, larger than the lower boundary and should not exceed: %v \n", MaxInt)
			os.Exit(1)
		}

	} else {
		fmt.Printf("Format error for input range: %v \n", err)
		os.Exit(1)
	}

	return l, u
}

// Function reads input for the algorithm flag and
// returns associated algorithm type.
// Exit if the input entered by the user is anything other than 1, 2 or 3.
func getAlgorithm(aPtr *int) (a int) {

	switch *aPtr {
	case algoBasic:
		a = algoBasic
	case algoEratos:
		a = algoEratos
	case algoSSE:
		a = algoSSE
	default:
		fmt.Printf("Algorithm type must be 1, 2 or 3. \n")
		os.Exit(1)
	}

	return a
}

// Function reads command line arguments and parses them.
// Then fills the struct through its pointer.
func parseFlags(inFlags *Flags) {

	var rangeStr string

	defStr := `Input range should be non-negative integers. 
	Ex: --range=1,100`
	flag.StringVar(&rangeStr, "range", "", defStr)

	defStr = `Specify generation algorithm. 
	1 for basic. 
	2 for Sieve of Eratosthenes.
	3 for Segmented SoE. 
	Ex: --algorithm=3.`
	algoPtr := flag.Int("algorithm", 2, defStr)

	defStr = `Measure wall clock time. 
	Ex: --measure=true. 
	Default is false.`
	measurePtr := flag.Bool("measure", false, defStr)

	defStr = `Print results to stdout. 
	Ex: --print=true. 
	Default is false.`
	printPtr := flag.Bool("print", false, defStr)

	defStr = `Validate results. 
	Ex: --validate=true. 
	Default is false.`
	validPtr := flag.Bool("validate", false, defStr)

	defStr = `Enable parallel execution. 
	Ex: --parallel=true. 
	Default is false.`
	parPtr := flag.Bool("parallel", false, defStr)

	flag.Parse()

	inFlags.lowerBound, inFlags.upperBound = getRange(rangeStr)
	inFlags.algorithm = getAlgorithm(algoPtr)
	inFlags.isPrint = *printPtr
	inFlags.measure = *measurePtr
	inFlags.validate = *validPtr
	inFlags.parallel = *parPtr
}
