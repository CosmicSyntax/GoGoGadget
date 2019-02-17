package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// mat... requires float64

	// VECTORS

	a := mat.NewVecDense(3, []float64{1, 2, 3})
	b := mat.NewVecDense(3, []float64{4, 5, 6})

	w := mat.NewVecDense(3, nil) // set the struct to nil b/c it's a pointer
	w.AddVec(a, b)

	x := mat.NewVecDense(1, nil) // receive a vector vector multication... needs a 1x1
	// a.MulVec(a.T(), b) this does not work b/c a is not 1 x 1 ... damn
	x.MulVec(a.T(), b)

	matPrint(x)

	matPrint(w)

	w.AddScaledVec(a, 2, b) // a + 2 * b

	b.AddVec(b, a) // b is overwritten... spaces memory

	matPrint(w)
	matPrint(b)

	// MATRIX

	m := mat.NewDense(3, 4,
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	)
	matPrint(m)

	/* prints...
	⎡1   2   3   4⎤
	⎢5   6   7   8⎥
	⎣9  10  11  12⎦
	*/

	// look up values... getter

	fmt.Println(m.At(0, 2)) // 3 (Col, Row)

	// change values... setter

	m.Set(0, 2, 5)
	//m.SetCol

	fmt.Println(m.At(0, 2)) // 3 (Col, Row)

	// Multiply

	m2 := mat.NewDense(3, 4, nil)
	m2.Add(m, m)

	d := mat.NewDense(3, 3, nil)
	d.Product(m, m2.T()) // m * m2'

	matPrint(d)
	/* checks out on octave!
	⎡ 92  168  264⎤
	⎢168  348  556⎥
	⎣264  556  892⎦
	*/

}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
