# genPrime

This small piece of code detects prime numbers in a given range. 
```
$ genPrime --range=20,40
Elapsed time: 7.316µs for range: [20,40] using algorithm: 2 
        Number of primes: 4 
        Largest prime: 37 
```

## Compile and install 
First you need to get the code preferably using `git clone`:
```
$ git clone https://github.com/fertinaz/genPrime.git
```
Then change directory and go inside the project folder:
```
$ cd genPrime
````
You can now type the following command -- Assuming that you have 
successfully installed `GO` in your environment:
```
$ go install
```
Command above compiles the source code, and puts binary executable 
under `$GOPATH/bin`. If `$GOPATH/bin` is added to your `$PATH`, then 
you should be able to run the command below successfully:
```
$ genPrime --help
```
This command prints usage information with explaining options and their
default values. 
```
Usage of genPrime:
  -algorithm int
        Specify generation algorithm. 
                1 for basic. 
                2 for Sieve of Eratosthenes.
                3 for Segmented SoE. 
                Ex: --algorithm=3. (default 2)
  -parallel
        Parallel execution. 
                Ex: --parallel=true. 
                Default is false.
  -print
        Print results to stdout. 
                Ex: --print=true. 
                Default is false.
  -range string
        Input range should be non-negative integers. 
                Ex: --range=1,100
  -validate
        Validate results. 
                Ex: --validate=true. 
                Default is false.
```
Alternatively, you can use `go build` which compiles the code and leaves the resulting binary executable in the project directory. If you don't want to use `$GOPATH/bin` for some reason, you can stick to `go build` rather than 
`go install`.

## Run
Sample usage in the shortest form:
```
$ genPrime --range=35,75 
Elapsed time: 6.316µs for range: [35,75] using algorithm: 2 
        Number of primes: 10 
        Largest prime: 73 
```
This command executes `genPrime` using the second algorithm which
is the Sieve of Eratosthenes. 

Flags can be specified both with a single dash `-` or double dash `--`. 
```
$ genPrime -range=10,30
Elapsed time: 5.19µs for range: [10,30] using algorithm: 2 
        Number of primes: 6 
        Largest prime: 29 
```

An execution with a complete set of flags:
```
$ genPrime -range=40,80 -algorithm=3 -print -validate --parallel=false
Elapsed time: 19.578µs for range: [40,80] using algorithm: 3 
         Number of primes: 10 
         Largest prime: 79 
Results: 
         List of prime numbers: [41 43 47 53 59 61 67 71 73 79] 
Validation: 
         List of prime numbers: [41 43 47 53 59 61 67 71 73 79] 
-- Results are correct!
```
This execution uses the third algorithm (Segmented Sieve of Eratosthenes) 
to find the prime numbers in the range between 40 and 80. 
Since `-print` option is provided it'll print the whole list of results. 
Those results are compared against the hardcoded list of primes because
`-validate` option is enabled. Parallel execution is disabled explicitly
which is the default setting anyway.

Note:
--parallel option doesn't work properly, so please ignore it at the moment.

## Tests
Additionally, to run some simple tests:
```
$ go test -v
```

## Results
Some results produced with a decent mobile workstation:

| n   | number of primes | Basic algorithm | Eratosthenes | Segmentation |
| --- | ---: | ---: | ---:| ---: |
| 2   |             25 |     0.000008786 |  0.000002744 |  0.000009336 | 
| 3   |            168 |     0.000434863 |  0.000028704 |  0.000025178 | 
| 4   |           1229 |     0.028150275 |  0.000085781 |  0.000124532 | 
| 5   |           9592 |     1.967347733 |  0.001020670 |  0.000962392 | 
| 6   |          78498 |   168.564263499 |  0.013824468 |  0.009132511 | 
| 7   |         664579 |              NA |  0.288736999 |  0.070968662 | 
| 8   |        5761455 |              NA |  6.574295088 |  0.668643781 | 
| 9   |       50847534 |              NA | 80.710604789 |  7.068193302 | 

Table presents wall clock times in seconds for a given range [1, N] 
where N = 10^n.
