package main

func checkPrime(num int) bool {

	var temp int

	// Iterate over every smaller integer and check remainder.
	// 0 remainder means number is divisible, so cannot be prime.
	for i := 3; i < num/2; i++ {
		temp = num % i
		if temp == 0 {
			return false
		}
	}

	return true
}

//	Function checks numbers between the [l, u] range
//	Input:
//		l: lower bound of the input range -- int
//		u: upper bound of the input range -- int
//	Return:
//		p: List of primes between input range
func basic(l, u int) Primes {

	var isPrime bool
	var p Primes

	// Adjust lower bound: 2 is the smallest prime number
	if l < 3 {
		p = append(p, 2)
		l = 3
	}

	// Adjust lower bound: even numbers cannot be prime,
	// so no need to iterate over them
	if l%2 == 0 {
		l++
	}

	for i := l; i <= u; i += 2 {
		isPrime = checkPrime(i)
		if isPrime {
			p = append(p, i)
			isPrime = false
		}
	}

	return p
}
