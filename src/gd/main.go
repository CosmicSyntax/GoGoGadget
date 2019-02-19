package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"unsafe"
)

type housing []int

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

		data["sqFeet"] = append(data["sqFeet"], var1)
		data["bedRoom"] = append(data["bedRoom"], var2)
		data["price"] = append(data["price"], var3)

	}

	var test int // for checking size of data

	for _, vars := range dataVar {
		fmt.Println(vars+":", data[vars])
		fmt.Println("Size on mem:", int(unsafe.Sizeof(test))*len(data[vars])) // size of each slice
	}

}
