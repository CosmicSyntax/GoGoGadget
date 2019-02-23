package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
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

	// Get size of training set
	length := len(data["sqFeet"])

	// Setting the learning rate
	var alpha float64 = .03

	// Theta...
	theta := mat.NewDense(3, 1, []float64{0, 0, 0}) // possible answer 139.21067X1 - 8738.01911X2 + 89597.90954
	MP(theta)

	// Training sets...
	//training := mat.NewDense(length, 3, nil)
	theta0 := make([]float64, length)

	for i, _ := range theta0 {
		theta0[i] = 1
	}
	// fmt.Println(theta0)

	// Feature scaling...
	// (xi - ui) / stdi
	// Calculate the mean
	var meanSq float64
	var meanBr float64

	for i, n := range data["sqFeet"] {
		meanSq += n
		if i == length-1 {
			meanSq = meanSq / float64(length)
		}
	}
	for i, n := range data["bedRoom"] {
		meanBr += n
		if i == length-1 {
			meanBr = meanBr / float64(length)
		}
	}

	var stdSq float64
	var stdBr float64

	for i := 0; i < length; i++ {
		// for Sq
		stdSq += math.Pow(data["sqFeet"][i]-meanSq, 2)
		// for Br
		stdBr += math.Pow(data["bedRoom"][i]-meanBr, 2)
	}

	stdSq = math.Sqrt(stdSq / float64(length-1))
	stdBr = math.Sqrt(stdBr / float64(length-1))

	// apply to data
	for i := 0; i < length; i++ {
		data["sqFeet"][i] = (data["sqFeet"][i] - meanSq) / stdSq
		data["bedRoom"][i] = (data["bedRoom"][i] - meanBr) / stdBr
	}

	// making the training set
	training := mat.NewDense(length, 3, nil)
	for i, _ := range theta0 {
		theta0[i] = 1
	}

	training.SetCol(0, theta0)
	training.SetCol(1, data["sqFeet"])
	training.SetCol(2, data["bedRoom"])
	MP(training)

	yHat := mat.NewDense(length, 1, data["price"])

	for i := 0; ; i++ {
		// Gradient
		grad := mat.NewDense(3, 1, nil)
		tXt := mat.NewDense(length, 1, nil)
		tXt.Product(training, theta)
		tXt.Sub(tXt, yHat)

		grad.Product(training.T(), tXt)
		grad.Scale(1/float64(length), grad)

		// apply alpha and change theta
		grad.Scale(alpha, grad)
		theta.Sub(theta, grad)

		MP(theta)

		time.Sleep(10 * time.Millisecond)
	}

}

func MP(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
