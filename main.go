package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type vec2 int32

func (v vec2) xy() (int16, int16) {
	y := int16(v % (1 << 16))
	x := int16((int32(v) - int32(y)) / (1 << 16))
	return x, y
}
func (v vec2) String() string {
	x, y := v.xy()
	return fmt.Sprintf("(%v, %v)", x, y)
}

func WriteJSON(positions []vec2, name string) {
	length := len(positions)
	xList := make([]int16, length)
	yList := make([]int16, length)
	for i, v := range positions {
		x, y := v.xy()
		xList[i] = x
		yList[i] = y
	}

	file, err := os.Create(fmt.Sprintf("%v.json", name))

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Println("Error closing file:", cerr)
		}
	}()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	dataObject := struct {
		X []int16 `json:"x"`
		Y []int16 `json:"y"`
	}{xList, yList}

	// Encode the data and write it to the file
	err = encoder.Encode(dataObject)

	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("JSON data written to output.json")

}

const right = vec2(1 << 16)
const left = -right
const up = vec2(1)
const down = -up

var unitVectors = []vec2{right, up, left, down}

func SimpleRandomWalk(n int) *[]vec2 {
	res := make([]vec2, n+1)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		res[i+1] = res[i] + unitVectors[r.Intn(4)]
	}

	return &res
}

func main() {

	n := 100
	var err error
	if len(os.Args) > 1 {
		if n, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatalf("Invalid walk length")
		}
	}
	procs := 10
	if len(os.Args) > 2 {
		if procs, err = strconv.Atoi(os.Args[2]); err != nil {
			log.Fatalf("Invalid process count")
		}
	}

	start := time.Now()
	pos := PooledMaplessSA(n, 30, procs)

	WriteJSON(*pos, fmt.Sprintf("output/sa%v", n))

	elapsed := time.Since(start)
	log.Printf("Search took %s", elapsed)
}
