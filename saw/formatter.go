package saw

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSON(positions []Vec2, name string) {
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
