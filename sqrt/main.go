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

// p = (x+1)*(y+s)
func equation_solution_check (p *big.Int, x_ []byte, y_ []byte) bool {
	x := binArr2BInt(x_)
	y := binArr2BInt(y_)
	
	x1 := new(big.Int).Add(x,big.NewInt(1))
	
	s := new(big.Int).Sqrt(p)
	
	y_s := new(big.Int).Add(y,s)
	
	x1ys := new(big.Int).Mul(x1,y_s)
	
	return p.Cmp(x1ys) == 0
}

// p && mask = (x+1)*(y+s) && mask
func equation_1 (p *big.Int, x_ []byte, y_ []byte) bool {
	x := binArr2BInt(x_)
	y := binArr2BInt(y_)
	
	x1 := new(big.Int).Add(x,big.NewInt(1))
	s := new(big.Int).Sqrt(p)
	y_s := new(big.Int).Add(y,s)
	
	x1ys := new(big.Int).Mul(x1,y_s)
	
	mask := bitMask(len(x_))
	
	p_mask := new(big.Int).And(mask, p)
	x1ys_mask := new(big.Int).And(mask, x1ys)

	return p_mask.Cmp(x1ys_mask) == 0
}

//  (p^2-s^2) && mask = [s^2*(x^2+2x)+2s(yx^2+2xy+y)+x^2y^2+2xy^2+y^2] && mask
func equation_2 (p *big.Int, x_ []byte, y_ []byte) bool {
	mask := bitMask(len(x_))

	x := binArr2BInt(x_)
	y := binArr2BInt(y_)

	s := new(big.Int).Sqrt(p)
	ss:= new(big.Int).Mul(s,s)
	pp:= new(big.Int).Mul(p,p)
	pp_ss:=new(big.Int).Sub(pp,ss)
	
	pp_ss_mask := new(big.Int).And(mask, pp_ss)
	
	
	xx:=new(big.Int).Mul(x,x)
	yy:=new(big.Int).Mul(y,y)
	
	x2:=new(big.Int).Mul(x,big.NewInt(2))

	f1:=new(big.Int).Add(xx,x2)
	ssf1:=new(big.Int).Mul(ss,f1)
	
	xy2:=new(big.Int).Mul(x2,y)
	xxy:=new(big.Int).Mul(xx,y)
	f2:=new(big.Int).Add(xxy,xy2)
	f2=f2.Add(f2,y)
	
	sf2:=new(big.Int).Mul(s,f2)
	s2f2:=new(big.Int).Mul(sf2,big.NewInt(2))
	
	xxyy:=new(big.Int).Mul(xx,yy)
	x2yy:=new(big.Int).Mul(x2,yy)
	
	sum:=new(big.Int).Add(ssf1,s2f2)
	sum=sum.Add(sum,xxyy)
	sum=sum.Add(sum,x2yy)
	sum=sum.Add(sum,yy)
	
	sum_mask := new(big.Int).And(mask, sum)
	
	fmt.Println(fmt.Sprintf("%b", mask));
	
	return pp_ss_mask.Cmp(sum_mask) == 0
}

func isCorrectEquationSystem(p *big.Int, first []byte, second []byte) (bool, bool) {
	
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

func bfs(p *big.Int, filter func (p *big.Int, first []byte, second []byte) (bool,bool), showSolution func (n1 *big.Int, n2 *big.Int)) {
	
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


				found, isCorrect := filter(p, BInt2binArr(first,n+1), BInt2binArr(second,n+1))
			
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
	var x, _ = new(big.Int).SetString("193", 10) //
	var y, _ = new(big.Int).SetString("173", 10) //
	p := big.NewInt(0).Mul(x, y)
	
	fmt.Println(fmt.Sprintf("x        :%32b=[%10X]=%v", x,x,x))
	fmt.Println(fmt.Sprintf("y        :%32b=[%10X]=%v", y,y,y))
	fmt.Println(fmt.Sprintf("p  =x*y  :%32b=[%10X]=%v", p,p,p))
	
	s := new(big.Int).Sqrt(p)
	
	fmt.Println(fmt.Sprintf("s        :%v", s))
	
	fmt.Println("-----------------------------------");
	
	bfs(p, isCorrectEquationSystem, showSolution)
}