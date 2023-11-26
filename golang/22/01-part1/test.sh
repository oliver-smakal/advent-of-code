#!/bin/bash

for infile in tests/*.in; do
    basename=${infile%.*}
    cat $infile | $1  > $basename.output
    diff $basename.output $basename.sol
done
