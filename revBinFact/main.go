package main

import (
	"fmt"
	"math/big"
)

func binArr2BInt(arr []byte) *big.Int {
	r := big.NewInt(0)
	for i:=len(arr)-1;i>=0 ;i-- {
		r.Mul(r,big.NewInt(2))
		r.Add(r,big.NewInt((int64)(arr[i])))
	}
	return r
}

const calculationBitsLimit = 3

func isCorrectTail(counter []byte) bool {
	fmt.Println(counter);
	if (len(counter) < calculationBitsLimit) {
		return true;
	}
	return false;
}

func binarySearch(filter func (counter []byte) bool ){
	counter := []byte{0}
	
	notFinished: for {
		isCorrect := filter(counter)

		if (isCorrect) {  			// correct, move to next register
			counter = append(counter,0)
		} else {					// incorrect, increment last one or return to previous register

			nextOne: for {
				item := counter[len(counter)-1]
				
				if (item == 3) {
					counter=counter[:len(counter)-1] // remove last one
				} else {
					item++
					counter[len(counter)-1] = item
					break nextOne
				}

				if (len(counter)==0) { break notFinished;}
			}
		}
		if (len(counter) == 0) {break}
	}
}

func main() {
	var x, _ = new(big.Int).SetString("173", 10)
	var y, _ = new(big.Int).SetString("197", 10)
	var p = big.NewInt(0).Mul(x, y)
	
	fmt.Println(fmt.Sprintf("x      :%b=%v", x,x));
	fmt.Println(fmt.Sprintf("y      :%b=%v", y,y));
	fmt.Println(fmt.Sprintf("product:%b=%v", p,p));
	fmt.Println("-----------------------------------");
	
	binarySearch(isCorrectTail);
	
	


	//aa := binArr2BInt(a)
	//bb := binArr2BInt(b)
	
	//fmt.Println(fmt.Sprintf("a      :%b=%v", aa,aa));
	//fmt.Println(fmt.Sprintf("b      :%b=%v", bb,bb));

	
			
}