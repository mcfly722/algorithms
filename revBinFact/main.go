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


func main() {
	var x, _ = new(big.Int).SetString("173", 10)
	var y, _ = new(big.Int).SetString("197", 10)
	var p = big.NewInt(0).Mul(x, y)
	
	fmt.Println(fmt.Sprintf("x      :%b=%v", x,x));
	fmt.Println(fmt.Sprintf("y      :%b=%v", y,y));
	fmt.Println(fmt.Sprintf("product:%b=%v", p,p));
	fmt.Println("-----------------------------------");
	
	a := []byte{1,0,1,0,0,1}
	b := []byte{0,0,1,1,1,1}

	aa := binArr2BInt(a)
	bb := binArr2BInt(b)
	
	fmt.Println(fmt.Sprintf("a      :%b=%v", aa,aa));
	fmt.Println(fmt.Sprintf("b      :%b=%v", bb,bb));

	
			
}