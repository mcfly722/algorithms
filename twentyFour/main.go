package main

import (
	"fmt"
	"math/big"
)

const searchDepthBits = 8

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


// p = (2m+1)*(2n+1)
func equation_solution_check (p *big.Int, m_ []byte, n_ []byte) bool {
	m := binArr2BInt(m_)
	n := binArr2BInt(n_)
	
	m2 := new(big.Int).Mul(m,big.NewInt(2))
	n2 := new(big.Int).Mul(n,big.NewInt(2))
	
	m2_1 := new(big.Int).Add(m2,big.NewInt(1))
	n2_1 := new(big.Int).Add(n2,big.NewInt(1))

	m2_1n2_1 := new(big.Int).Mul(m2_1,n2_1)
	
	return p.Cmp(m2_1n2_1) == 0
}

// p && mask = (2m+1)*(2n+1) && mask
func equation_1 (p *big.Int, m_ []byte, n_ []byte) bool {
	m := binArr2BInt(m_)
	n := binArr2BInt(n_)
	
	m2 := new(big.Int).Mul(m,big.NewInt(2))
	n2 := new(big.Int).Mul(n,big.NewInt(2))
	
	m2_1 := new(big.Int).Add(m2,big.NewInt(1))
	n2_1 := new(big.Int).Add(n2,big.NewInt(1))

	m2_1n2_1 := new(big.Int).Mul(m2_1,n2_1)
	
	mask := bitMask(len(m_))
	
	p_mask := new(big.Int).And(mask, p)
	m2_1n2_1_mask := new(big.Int).And(mask, m2_1n2_1)

	return p_mask.Cmp(m2_1n2_1_mask) == 0
}

//  ((p-1)/2) && mask = (m+n+2mn) && mask
func equation_2 (p *big.Int, m_ []byte, n_ []byte) bool {
	m := binArr2BInt(m_)
	n := binArr2BInt(n_)

	p_1 := new(big.Int).Sub(p,big.NewInt(1))
	d := new(big.Int).Div(p_1,big.NewInt(2))

	m_n := new(big.Int).Add(m,n)
	mn := new(big.Int).Mul(m,n)
	mn2 := new(big.Int).Mul(mn,big.NewInt(2))
	m_n_2mn :=new(big.Int).Add(m_n,mn2)
	
	mask := bitMask(len(m_))

	d_mask := new(big.Int).And(mask, d)
	m_n_2mn_mask := new(big.Int).And(mask, m_n_2mn)
	
	return d_mask.Cmp(m_n_2mn_mask) == 0
}

func product(numbers []int64) *big.Int {
	if len(numbers)==0 {
		return big.NewInt(0)
	}
	
	product := big.NewInt(1)
	
	for _, number := range numbers {
		product = product.Mul(product,big.NewInt(number))
	}
	
	return product 
}


func isCorrectEquationSystem(p *big.Int, dividers []int64, square int64, first []byte, second []byte) (bool, bool) {
	
	if (len(first) > searchDepthBits){
		return false, false;
	}

	eq_1 := equation_1(p, first, second)
	eq_2 := equation_2(p, first, second)

	fmt.Println(fmt.Sprintf("%s %s %5t %5t", bin2str(first),bin2str(second), eq_1, eq_2));
//	fmt.Println("-----------------------------------------------------------------")
	
	
	if eq_1 && eq_2 {
		return equation_solution_check(p, first,second), true;
	}
	return false, false;
}


type pair struct {
	first *big.Int
	second *big.Int
}

func showSolution(n1 *big.Int, n2 *big.Int) {
	fmt.Println(fmt.Sprintf("\nSOLUTION:%v,%v",n1,n2));
}

func getSquareAndDividers(p *big.Int) ([]int64, int64) {
	dividers := []int64{}
	square := int64(1)
	
	squares: for j:=1;j<1000;j++ {
		p_jj := new(big.Int).Sub(p, big.NewInt(int64(j*j)))
		
		if p_jj.Cmp(big.NewInt(0)) < 0 {
			break squares
		}
		
		tmp_dividers := []int64{1}
		
		for i:=2;i<1000;i++ {
			if big.NewInt(int64(i)).ProbablyPrime(1) {
				mod :=new(big.Int).Mod(p_jj,big.NewInt(int64(i)))
				if mod.Cmp(big.NewInt(0))==0 {
					tmp_dividers = append(tmp_dividers, int64(i))
				}
			}
		}

		if len(tmp_dividers) > len(dividers) {
			square = int64(j)
			dividers = make([]int64,len(tmp_dividers))
			copy (dividers,tmp_dividers)
			//fmt.Println(fmt.Sprintf("p-sqrt=%v  div=%v       final=%v", p_jj,tmp_dividers, dividers))
		}
	}

	return dividers, square
}

func bfs(p *big.Int, dividers []int64,square int64, filter func (p *big.Int, dividers []int64,square int64, first []byte, second []byte) (bool,bool), showSolution func (n1 *big.Int, n2 *big.Int)) {
	
	layer := []pair{
		pair{
			first: big.NewInt(0),
			second: big.NewInt(0),},
	}
	
	search: for n:=0;n<searchDepthBits;n++ {

		var newLayer = []pair{};

		fmt.Println(fmt.Sprintf("LAYER %v", n))		
		
		for _,pair_ := range layer {

			for i:=0;i<4;i++ {
				bit1 := i & 1
				bit2 := (i/2) & 1
			
				first := new(big.Int).Exp(big.NewInt(2),big.NewInt(int64(n)),nil)
				first = first.Mul(first,big.NewInt(int64(bit1)))
				first.Add(first, pair_.first)

				second := new(big.Int).Exp(big.NewInt(2),big.NewInt(int64(n)),nil)
				second = second.Mul(second,big.NewInt(int64(bit2)))
				second.Add(second, pair_.second)


				found, isCorrect := filter(p, dividers, square, BInt2binArr(first,n+1), BInt2binArr(second,n+1))
			
				if isCorrect {

					if found {
						showSolution(first, second)
						break search
					
					} else {
						newLayer = append(newLayer, pair {
							first: first,
							second: second})
					}
				}
			} 
		}

		fmt.Println(fmt.Sprintf("            COUNT: %v",len(newLayer)))
		
		layer = newLayer;
	}
}

func main() {
	var x, _ = new(big.Int).SetString("2333", 10) //
	var y, _ = new(big.Int).SetString("173", 10) //
	p := big.NewInt(0).Mul(x, y)
	
	fmt.Println(fmt.Sprintf("x        :%32b=[%10X]=%v", x,x,x))
	fmt.Println(fmt.Sprintf("y        :%32b=[%10X]=%v", y,y,y))
	fmt.Println(fmt.Sprintf("p  =x*y  :%32b=[%10X]=%v", p,p,p))
	
	dividers, square := getSquareAndDividers(p)
	
	p_s := new(big.Int).Sub(p,big.NewInt(square*square))
	fmt.Println(fmt.Sprintf("square   :%v", square))
	fmt.Println(fmt.Sprintf("p-square :%v", p_s))
	fmt.Println(fmt.Sprintf("dividers :%v", dividers))
	fmt.Println(fmt.Sprintf("div.prod.:%v", product(dividers)))
	
	fmt.Println("-----------------------------------");
	
//	bfs(p, dividers, square, isCorrectEquationSystem, showSolution)
}