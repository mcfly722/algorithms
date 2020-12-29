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

/*
func equation_1 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	b := binArr2BInt(second)


	x := new(big.Int)
	x.Set(a)
	x = x.Add(x,b)
	
	y := new(big.Int)
	y.Set(b)
	y = y.Sub(y,a)
		
	xy := new(big.Int).Set(x)
	xy = xy.Mul(xy,y)
	
	mask := bitMask(len(first))

	xy_f := new(big.Int).Set(xy).And(mask,xy)
	t_f := new(big.Int).Set(p).And(mask,p)
	
	areEqual := xy_f.Cmp(t_f) == 0
	
	
	fmt.Println(fmt.Sprintf("e_1:[%v,%v](%v,%v)->(%v,%v)->%v && %b->  %v[%b] ?= %v[%b]  => %t",
		bin2str(first),bin2str(second),
		a,b,
		x,y,
		xy,mask,
		xy_f,xy_f,
		t_f,t_f,
		areEqual));
	
	
	return areEqual
}

func equation_2 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	b := binArr2BInt(second)

	aa := new(big.Int).Set(a)
	aa = aa.Mul(aa,aa)
	
	bb := new(big.Int).Set(b)
	bb = bb.Mul(bb,bb)
	
	mask := bitMask(len(first))
	
	aa_c := new(big.Int).Set(aa).Add(aa,p)
	
	aa_c_f := new(big.Int).Set(aa_c).And(mask,aa_c)
	
	bb_f := new(big.Int).Set(bb).And(mask,bb)

	areEqual := aa_c_f.Cmp(bb_f) == 0
	
	fmt.Println(fmt.Sprintf("e_2:[%v,%v](%v,%v)^2->%v,%v->+%v[%b] = %v[%b] && [%b]-> %v[%b] ?= %v[%b]  => %t",
		bin2str(first),bin2str(second),
		a,b,
		aa,bb,
		p,p,
		aa_c,aa_c,
		mask,
		aa_c_f,aa_c_f,
		bb_f,bb_f,
		areEqual));
		
	return areEqual
}
*/
/*
func equation_3 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	b := binArr2BInt(second)

	aa := new(big.Int).Set(a)
	aa = aa.Mul(aa,a)

	sum :=aa.Add(aa,p)

	bb := new(big.Int).Set(b)
	bb = bb.Mul(bb,b)

	//fmt.Println(fmt.Sprintf("%v, %v",bin2str(first),bin2str(second)));
	return sum.Cmp(bb) == 0
}

*/


/*
func dfs(filter func (first []byte, second []byte) (bool,bool),showSolution func (n1 *big.Int, n2 *big.Int)) {
	
	counter := []byte{0}
	
	notFinished: for {
		first:=[]byte{};
		second :=[]byte{};

		for _, value :=range counter {
			first = append(first,value & 1)
			second = append(second,(value/2) & 1)
		}

		founded, isCorrect := filter(first, second)
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

*/

/*
func equation_4 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	d := binArr2BInt(second)

	aa := new(big.Int).Set(a)
	aa = aa.Mul(aa,a)

	sum :=aa.Add(aa,p)
	
	d_a := d.Add(d,a)
	d_ad_a := d_a.Mul(d_a,d_a)
	
	return sum.Cmp(d_ad_a) == 0
}

func equation_5 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	d := binArr2BInt(second)

	aa := new(big.Int).Set(a)
	aa = aa.Mul(aa,a)

	sum :=aa.Add(aa,p)
	
	d_a := d.Add(d,a)
	d_ad_a := d_a.Mul(d_a,d_a)	
	
	mask := bitMask(len(first))
	
	sum_f := new(big.Int).Set(sum).And(mask,sum)
	d_ad_a_f := new(big.Int).Set(d_ad_a).And(mask,d_ad_a)
	
	return sum_f.Cmp(d_ad_a_f) == 0
}

*/





// p?=m^2-n^2
func equation_solution_check (p *big.Int, m_ []byte, n_ []byte) bool {
	m := binArr2BInt(m_)
	n := binArr2BInt(n_)

	mm := m.Mul(m,m)
	nn := n.Mul(n,n)
	nn_p := nn.Add(nn,p)
	
	return nn_p.Cmp(mm) == 0
}

// (p + n^2) && mask =m^2 && mask
func equation_1 (p *big.Int, m_ []byte, n_ []byte) bool {
	m := binArr2BInt(m_)
	n := binArr2BInt(n_)

	mm := new(big.Int).Mul(m,m)
	nn := new(big.Int).Mul(n,n)

	p_nn := new(big.Int).Add(p, nn)
	
	mask := bitMask(len(m_))

	p_nn_m := new(big.Int).And(mask,p_nn)
	mm_m :=new(big.Int).And(mask,mm)

	if binArr2BInt(m_).Cmp(big.NewInt(57))==0 && binArr2BInt(n_).Cmp(big.NewInt(244))==0 {
		fmt.Println(fmt.Sprintf("m      =%3b n   =%3b", m,n));
		fmt.Println(fmt.Sprintf("mm     =%3b nn  =%3b", mm,nn));
		fmt.Println(fmt.Sprintf("p_nn   =%3b", p_nn));
		fmt.Println(fmt.Sprintf("mask   =%3b", mask));
		fmt.Println(fmt.Sprintf("p_nn_m =%3b", p_nn_m));
		fmt.Println(fmt.Sprintf("mm_m   =%3b", mm_m));
	}

	return p_nn_m.Cmp(mm_m) == 0
}

//  (p ^2+4*m^2*n^2) && mask=(m^2+n^2)^2 && mask
func equation_2 (p *big.Int, m_ []byte, n_ []byte) bool {
	m := binArr2BInt(m_)
	n := binArr2BInt(n_)

	mm := new(big.Int).Mul(m,m)
	nn := new(big.Int).Mul(n,n)

	mmnn := new(big.Int).Mul(mm,nn)
	mmnn4:= new(big.Int).Mul(mmnn,big.NewInt(4))

	pp := new(big.Int).Mul(p,p)
	pp_mmnn4 :=new(big.Int).Add(pp, mmnn4)

	mask := bitMask(len(m_))
	pp_mmnn4_m := new(big.Int).And(mask,pp_mmnn4)
	
	mm_nn :=new(big.Int).Add(mm,nn)
	mm_nnmm_nn :=new(big.Int).Mul(mm_nn,mm_nn)
	mm_nnmm_nn_m :=new(big.Int).And(mask,mm_nnmm_nn)


	if binArr2BInt(m_).Cmp(big.NewInt(57))==0 && binArr2BInt(n_).Cmp(big.NewInt(244))==0 {
		fmt.Println(fmt.Sprintf("m           =%3b n =%3b", m,n));
		fmt.Println(fmt.Sprintf("mm          =%3b nn=%3b", mm,nn));
		fmt.Println(fmt.Sprintf("mmnn        =%3b", mmnn));
		fmt.Println(fmt.Sprintf("mmnn4       =%3b", mmnn4));
		fmt.Println(fmt.Sprintf("pp          =%3b", pp));
		fmt.Println(fmt.Sprintf("pp_mmnn4    =%3b", pp_mmnn4));
		fmt.Println(fmt.Sprintf("mask        =%3b", mask));
		fmt.Println(fmt.Sprintf("pp_mmnn4_m  =%3b", pp_mmnn4_m));
		fmt.Println(fmt.Sprintf("mm_nn       =%3b", mm_nn));
		fmt.Println(fmt.Sprintf("mm_nnmm_nn  =%3b", mm_nnmm_nn));
		fmt.Println(fmt.Sprintf("mm_nnmm_nn_m=%3b", mm_nnmm_nn_m));
	}
	
	return pp_mmnn4_m.Cmp(mm_nnmm_nn_m) == 0
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

func showSolution(n1 *big.Int, n2 *big.Int) {
	fmt.Println(fmt.Sprintf("\nSOLUTION:%v,%v",n1,n2));
}

func main() {
	var x, _ = new(big.Int).SetString("197", 10) //
	var y, _ = new(big.Int).SetString("173", 10) //
	p := big.NewInt(0).Mul(x, y)

	m := new(big.Int).Set(x).Add(x,y)
	m = m.Div(m,big.NewInt(2))
	
	n := new(big.Int).Set(x).Sub(x,y)
	n = n.Div(n,big.NewInt(2))
	
	fmt.Println(fmt.Sprintf("x  :%32b=[%10X]=%v", x,x,x));
	fmt.Println(fmt.Sprintf("y  :%32b=[%10X]=%v", y,y,y));
	fmt.Println(fmt.Sprintf("x*y:%32b=[%10X]=%v", p,p,p));

	fmt.Println(fmt.Sprintf("m  :%32b=[%10X]=%v", m,m,m));
	fmt.Println(fmt.Sprintf("n  :%32b=[%10X]=%v", n,n,n));
	
	
	
	fmt.Println("-----------------------------------");
	
	bfs(p, isCorrectEquationSystem,showSolution)
}