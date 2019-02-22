package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
	"unsafe"

	"gonum.org/v1/gonum/mat"
)

type housing []float64

var data map[string]housing

var dataVar = []string{"sqFeet", "bedRoom", "price"}

func main() {

	// Open the text file
	csvF, _ := os.Open("./ex1data2.csv")
	defer csvF.Close()
	buf := bufio.NewReader(csvF) // reading in 16 bytes
	r := csv.NewReader(buf)

	// Read file...

	data = make(map[string]housing) // initialize the map

	for _, vars := range dataVar {
		data[vars] = housing{}
	}

	for {

		line, error := r.Read() // built in for loop for buf
		if error == io.EOF {
			break
		} else if error != nil {
			panic(error)
		}

		var1, _ := strconv.Atoi(line[0])
		var2, _ := strconv.Atoi(line[1])
		var3, _ := strconv.Atoi(line[2])

		data["sqFeet"] = append(data["sqFeet"], float64(var1))
		data["bedRoom"] = append(data["bedRoom"], float64(var2))
		data["price"] = append(data["price"], float64(var3))

	}

	var test float64 // for checking size of data

	for _, vars := range dataVar {
		fmt.Println(vars+":", data[vars])
		fmt.Println("Size on mem:", int(unsafe.Sizeof(test))*len(data[vars])) // size of each slice
	}

	// Linear alebgra set up...
	length := len(data["sqFeet"])

	// Setting the learning rate

	alpha := float64(0.00000001)

	// Theta...
	theta := mat.NewDense(3, 1, []float64{0, 0, 0}) // possible answer 139.21067X1 - 8738.01911X2 + 89597.90954
	MP(theta)

	// Training sets...
	training := mat.NewDense(length, 3, nil)
	theta0 := make([]float64, length)

	for i, _ := range theta0 {
		theta0[i] = 1
	}
	// fmt.Println(theta0)

	training.SetCol(0, theta0)
	training.SetCol(1, data["sqFeet"])
	training.SetCol(2, data["bedRoom"])
	MP(training)

	// Vector Prices
	yHat := mat.NewDense(length, 1, data["price"])
	MP(yHat)

	// ---- START MATH ---- //

	for {
		TT := mat.NewDense(length, 1, nil)
		TT.Product(training, theta)
		TT.Sub(TT, yHat)

		DT := mat.NewDense(3, 1, nil)
		DT.Product(training.T(), TT)

		DT.Scale(1/float64(length), DT)
		DT.Scale(alpha, DT)

		theta.Sub(theta, DT)

		// MP(theta)

		// calculate cost F(x)
		costFunc := mat.NewDense(1, 1, nil)
		costFunc.Product(TT.T(), TT)
		costFunc.Scale(1/float64(length), costFunc)

		fmt.Println(costFunc.At(0, 0))

		time.Sleep(10 * time.Millisecond)
	}

}

func MP(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
