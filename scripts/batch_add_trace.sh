#!/bin/bash 

files=$(find . -name "*.go" -not -name "*_test.go" -print)
echo $files

for f in $files
do 
		gen -w $f
done
