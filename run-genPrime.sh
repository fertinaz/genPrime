#!/bin/bash

### Go installation -- change if needed
# GOROOT=/usr/local/go
# GOBIN=${GOROOT}/bin

echo -e "Testing genPrime: " > test4.txt

for n in `seq 2 6`;
do
    echo -e $((10 ** $n))
    N=$((10 ** $n))
    go run *.go --range=1,$N --algorithm=1 --print=false >> test4.txt
    go run *.go --range=1,$N --algorithm=2 --print=false >> test4.txt
    go run *.go --range=1,$N --algorithm=3 --print=false >> test4.txt
done
