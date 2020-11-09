package main

import (
	"fmt"
	"math/big"
)
	
func main() {
	var x, _ = new(big.Int).SetString("173", 10)
	var y, _ = new(big.Int).SetString("197", 10)
	var p = new(big.Int).Set(x)
	p.Mul(x,y)
	
	fmt.Println(fmt.Sprintf("x      :%b=%v", x,x));
	fmt.Println(fmt.Sprintf("y      :%b=%v", y,y));
	fmt.Println(fmt.Sprintf("product:%b=%v", p,p));
	
	
}