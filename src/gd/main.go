package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type house struct {
	sqFeet  int
	bedRoom int
	price   int
}

func main() {
	// Open the text file
	csvF, _ := os.Open("./ex1data2.csv")
	defer csvF.Close()
	buf := bufio.NewReader(csvF) // reading in 16 bytes
	r := csv.NewReader(buf)

	// Read file...

	var housing []house

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

		housing = append(housing, house{
			sqFeet:  var1,
			bedRoom: var2,
			price:   var3,
		})
	}

	fmt.Println(housing)

}
