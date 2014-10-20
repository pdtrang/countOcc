package main

import (
	"fmt"
	"flag"
	"github.com/pdtrang/countOcc"
	"github.com/vtphan/fmi"
	"time"
	
)

var idx fmi.Index
var count_10 = 0
var count_20 = 0
var count_40 = 0
var count = 0

func update_count(k int){

	if (k == 10){
		count_10 = count_10 + 1
	}else{
		if (k == 20){
			count_20 = count_20 + 1
		}else{
			if (k == 40){
				count_40 = count_40 + 1
			}else{
				count = count + 1
			}
			
		}
	}
}

func search(sequence []byte, i int, k int){
	
	var arr = idx.Search(sequence[i:i+k])	
	
	if (len(arr) == 1){	
		update_count(k)
	}else{
		//occ > 1			
		if((2*k <= 40) && (i+2*k<len(sequence)) ){
			search(sequence, i, 2*k)
		}				
	}

}

func countO(sequence []byte, k int) {
	
	fmt.Println("Begin counting...")
	
	L := len(sequence)
	
	if(k<=40){
		for i := 0; i < L-k+1; i++ {
			search(sequence, i, k)		
		}
	}
	
	fmt.Println("Finish counting")

}

func main(){
	var genome_file = flag.String("g", "", "reference genome file")
	var index_file = flag.String("i", "", "index file of genome")
	
	flag.Parse()

	var k = 10

	s := time.Now()

	fmt.Println("Read FASTA")
    sequence := countOcc.ReadFASTA(*genome_file)

    //fmt.Println(string(sequence))
    fmt.Println(len(sequence))
    
	idx = *fmi.New(*genome_file)
	idx.Save(*index_file)
	//fmt.Println(idx)

	fmt.Println("Finish indexing genome...")

	countO(sequence, k)

	fmt.Println("\nNumber of occ = 1  ")
	fmt.Println("10-mer ", count_10, "\n20-mer ", count_20, "\n40-mer ", count_40, "\nn-mer ", count)

	end := time.Now()
	elapse := end.Sub(s)
	fmt.Println("Running time = ", elapse)
}
