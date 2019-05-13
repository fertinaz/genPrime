# genPrime

This piece of code is written for prime number detection in a given range. 

To compile and run -- assuming that you have successfully installed GO:
```
go run *.go --range=1,100 --algorithm=1
```
You can also use following command line flags for a more verbose execution:
```
go run *.go --range=1,100 --algorithm=1 --print=true --validate=true --parallel=false
```
This should also help:
```
go run *.go -h
```

You can also use makefile for building this code in the project directory:
```
make
./genPrime --range=1,200 --algorithm=3 --print=true --validate=false --parallel=false
```

Note:
--parallel option doesn't work properly, so you can ignore it at the moment.

Here are some results produced with a decent mobile workstation:

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

Table represents wall clock times in s. for a given range [1, N] where N = 10^n
