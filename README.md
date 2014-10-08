countOcc
========

Count how many 10-mer, 20-mer (extends from 10-mers which occur more than once), 40-mer (extends from 20-mers which occur more than once) occur once in a sequence

Input: FASTA file.

Output: number of 10-mer, 20-mer, 40-mer occur once.

Install: go get github.com/pdtrang/countOcc

go run main.go -g sequence.fasta -i indexfile
