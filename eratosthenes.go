package main

/*
	Function Sieve of Eratosthenes
	Input:
		l: lower bound of the input range -- int
		u: upper bound of the input range -- int
	Return:
		p: List of primes between input range
*/
func eratosthenes(l, u int) []int {

	var p []int
	pSet := make([]bool, u+1)

	for i := 2; i <= u; i++ {
		pSet[i] = true
	}

	for i := 2; i*i <= u; i++ {
		for j := i * 2; j <= u; j += i {
			pSet[j] = false
		}
	}

	for i := 2; i <= u; i++ {
		if pSet[i] {
			if i >= l {
				p = append(p, i)
			}
		}
	}

	return p
}
