package main

import (
	"fmt"
	"math/big"
	"math"
)

func Max(x, y int64) int64 {
    if x < y {
        return y
    }
    return x
}

func getPrimes(start int64, end int64) []int64 {
	dividers := []int64{}
	
	for i:=start;i<end;i++ {
		if big.NewInt(int64(i)).ProbablyPrime(1) {
			dividers=append(dividers, int64(i))
		}
	}

	return dividers
}

func getBottomSquares(p int64) []int64 {
	squares := []int64{}
	
	tmp := p
	
	for {
		s := int64( math.Floor(math.Sqrt( float64(tmp) )))
		squares=append(squares,s)
		tmp = tmp - s*s
		if tmp==0 {
			break
		}
	}

	return squares
}

func getTopSquares(p int64) []int64 {
	squares := []int64{}
	
	f := int64(math.Ceil(math.Sqrt(float64(p))))
	
	squares = append(squares,f)
	
	tmp := f*f-p
	
	for {
		s := int64( math.Floor(math.Sqrt( float64(tmp) )))

		if s>0 {
			squares=append(squares,s)
		}
		
		tmp = tmp-s*s


		if tmp == 0 {
			break
		}
		

	}

	return squares
}

func getBothSquares(p int64) ([]int64,[]int64) {
	first := []int64{}
	second := []int64{}
	
	t := int64(0)
	m := int64(0)
	n := int64(0)
	
	for {
		m = int64(math.Ceil(math.Sqrt(float64(p-t))))
		tt := t + m * m

		if m > 0 {
			first = append(first,  tt)
		} else {
			break
		}
		
	
		n = int64(math.Ceil(math.Sqrt(float64(tt - p))))
		t = tt - n*n
		
		if n > 0 {
			second = append(second, t)
		} else {
			break
		}

		//fmt.Println(fmt.Sprintf("%v,%v",m,n))

		if m == 2 && n == 2 {
			break
		}
	}
	
	return first, second
}



func main() {

	//primes1 := getPrimes(1100,1200)
	//primes2 := getPrimes(7410,7820)
	primes1 := getPrimes(1,100)
	primes2 := getPrimes(1,100)
	
	for _,i :=range primes1 {
		for _,j :=range primes2 {
			p := j*i
			
			a := (i+j) / 2
			b := Max(i,j)-a
			
			fmt.Println(fmt.Sprintf("%4v * %4v = %8v => [%4v]-[%4v] : %v = %v - %v",i,j, p, a, b, getBottomSquares(p),getBottomSquares(a*a), getBottomSquares(b*b)))
			
			//fmt.Print(fmt.Sprintf("%4v * %4v = %8v => [%4v]-[%4v] ",i,j, p, a, b))
			//first,second := getBothSquares(p)
			
//			fmt.Println(fmt.Sprintf(" = %v - %v",first,second))
			
		}
	
	
	}
}