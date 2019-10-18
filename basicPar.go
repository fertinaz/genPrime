package main

func worker(jobs <-chan int, results chan<- Primes, d int) {
	for n := range jobs {
		results <- basicPar(n, d)
	}
}

// 	Concurrent version of the basic algorithm
// 	Input:
//      n: job (sort of) rank from the worker pool
// 	Return:
// 		p: list of primes in the certain range of this job
func basicPar(n, d int) Primes {

	var p Primes
	var isPrime bool

	// Set lower / upper bounds for each job
	// in the worker pool
	l := inFlags.lowerBound + (n * d)
	u := (l + d) - 1

	if l < 3 {
		p = append(p, 2)
		l = 3
	}

	if l%2 == 0 {
		l++
	}

	for k := l; k <= u; k += 2 {
		isPrime = true
		for j := 3; j < l/2; j += 2 {
			if (k % j) == 0 {
				isPrime = false
				break
			}
		}
		if isPrime == true {
			p = append(p, k)
		}
	}

	return p
}
