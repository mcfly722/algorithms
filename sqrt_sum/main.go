package main

import (
	"fmt"
	"math/big"
)

func bin2str(s []byte) string {
	r := ""
	for i := 0; i<len(s); i++ {
		r=r+string(48+s[len(s)-1-i])
	}
	return r
}

func binArr2BInt(arr []byte) *big.Int {
	r := big.NewInt(0)
	for i:=len(arr)-1;i>=0 ;i-- {
		r.Mul(r,big.NewInt(2))
		r.Add(r,big.NewInt((int64)(arr[i])))
	}
	return r
}

func BInt2binArr(number *big.Int,l int) []byte {
	r := []byte{}
	current := new(big.Int).Set(number)
	
	for i:=0;i<l;i++{
		_,m :=current.DivMod(current,big.NewInt(2),big.NewInt(2))
		r = append(r,(uint8)(m.Int64()))
	}

	return r
}

func bitMask(length int) *big.Int {
	r := big.NewInt(0)
	for i:=0; i<length ;i++ {
		r=r.Mul(r,big.NewInt(2))
		r=r.Add(r,big.NewInt(1))
	}
	return r;
}


// p?=m^2+n^2
func equation_solution_check (p *big.Int, m_ []byte, n_ []byte) bool {
	m := binArr2BInt(m_)
	n := binArr2BInt(n_)

	mm := new(big.Int).Mul(m,m)
	nn := new(big.Int).Mul(n,n)
	
	nn_mm := new(big.Int).Add(nn,mm)
		
	return p.Cmp(nn_mm) == 0
}

// p && mask = (m^2+n^2) && mask
func equation_1 (p *big.Int, m_ []byte, n_ []byte) bool {
	m := binArr2BInt(m_)
	n := binArr2BInt(n_)

	mm := new(big.Int).Mul(m,m)
	nn := new(big.Int).Mul(n,n)
	
	nn_mm := new(big.Int).Add(nn,mm)
	
	mask := bitMask(len(m_))


	p_m := new(big.Int).And(mask,p)
	nn_mm_m :=new(big.Int).And(mask,nn_mm)

/*
	if binArr2BInt(m_).Cmp(big.NewInt(57))==0 && binArr2BInt(n_).Cmp(big.NewInt(244))==0 {
		fmt.Println(fmt.Sprintf("m      =%3b n   =%3b", m,n));
		fmt.Println(fmt.Sprintf("mm     =%3b nn  =%3b", mm,nn));
		fmt.Println(fmt.Sprintf("p_nn   =%3b", p_nn));
		fmt.Println(fmt.Sprintf("mask   =%3b", mask));
		fmt.Println(fmt.Sprintf("p_nn_m =%3b", p_nn_m));
		fmt.Println(fmt.Sprintf("mm_m   =%3b", mm_m));
	}
*/
	return p_m.Cmp(nn_mm_m) == 0
}

func isCorrectEquationSystem(p *big.Int, first []byte, second []byte) (bool, bool) {
	
	if (len(first) > p.BitLen()){
		return false, false;
	}

	eq_1 := equation_1(p, first, second)

	
	if eq_1 {
		//fmt.Println(fmt.Sprintf("%s %s", bin2str(first), bin2str(second)));
		
		return equation_solution_check(p, first,second), true;
	}
	return false, false;
}

type pair struct {
	first *big.Int
	second *big.Int
}


func dfs(p *big.Int, filter func (p *big.Int, first []byte, second []byte) (bool,bool), showSolution func (n1 *big.Int, n2 *big.Int)) {


	counter := []byte{0}
	
	notFinished: for {
		first:=[]byte{};
		second :=[]byte{};

		for _, value :=range counter {
			first = append(first,value & 1)
			second = append(second,(value/2) & 1)
		}

		founded, isCorrect := filter(p , first, second)
		if founded {
		
			solution1:= new(big.Int).Set(binArr2BInt(first))
			solution2:= new(big.Int).Set(binArr2BInt(second))
			
			showSolution (solution1, solution2)
		}

		if isCorrect {  			// correct, move to next register
			counter = append(counter,0)
		} else {					// incorrect, increment last one or return to previous register
			nextOne: for {
				item := counter[len(counter)-1]
				if (item > 2) {
					counter=counter[:len(counter)-1] // remove last one
				} else {
					item++
					counter[len(counter)-1] = item
					break nextOne
				}

				if (len(counter)==0) {
					break notFinished
				}
				
				
			}
		}
		if (len(counter) == 0) {break}
	}
}

func showSolution(n1 *big.Int, n2 *big.Int) {
	fmt.Println(fmt.Sprintf("\nSOLUTION:%v,%v",n1,n2));
}

func main() {
	var x, _ = new(big.Int).SetString("115249", 10) // 197
	var y, _ = new(big.Int).SetString("23497", 10)  // 173
	p := big.NewInt(0).Mul(x, y)

	m := new(big.Int).Set(x).Add(x,y)
	m = m.Div(m,big.NewInt(2))
	
	n := new(big.Int).Set(x).Sub(x,y)
	n = n.Div(n,big.NewInt(2))
	
	fmt.Println(fmt.Sprintf("x  :%64b=[%10X]=%v", x,x,x));
	fmt.Println(fmt.Sprintf("y  :%64b=[%10X]=%v", y,y,y));
	fmt.Println(fmt.Sprintf("x*y:%64b=[%10X]=%v", p,p,p));

	fmt.Println(fmt.Sprintf("m  :%64b=[%10X]=%v", m,m,m));
	fmt.Println(fmt.Sprintf("n  :%64b=[%10X]=%v", n,n,n));
	
	
	
	fmt.Println("-----------------------------------");
	
	dfs(p, isCorrectEquationSystem,showSolution)
}