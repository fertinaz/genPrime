#!/bin/bash

### Go installation -- change if needed
# GOROOT=/usr/local/go
# GOBIN=${GOROOT}/bin

### Project directory
# base=$HOME/dev/go
# case=genPrime

### Source files
# src=genPrime.go


echo -e "Testing genPrime: " > test.txt

for bound in `seq 2 6`;
do
    echo -e $((10 ** $bound))
    upper=$((10 ** $bound))
    go run *.go --range=1,$upper --algorithm=1 --print=false >> test.txt
    go run *.go --range=1,$upper --algorithm=2 --print=false >> test.txt
    go run *.go --range=1,$upper --algorithm=3 --print=false >> test.txt
done
