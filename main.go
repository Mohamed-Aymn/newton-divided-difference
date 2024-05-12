package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	x = 1.5
)

type DataPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func main() {
	/**
	**
	** Read Data
	**
	**/
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the JSON data into a struct
	var data []DataPoint
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// number of points detection
	n := len(data)

	/**
	**
	** calculate data
	**
	**/
	var calculations [][]float64
	for degree := 0; degree < n-1; degree++ {
		calculations = append(calculations, make([]float64, n-degree))
		dividedDiff(degree, data, &calculations, &n)
	}

	/**
	**
	** calculate final result
	**
	**/
	result := calculate(calculations, data, &n)
	fmt.Println("Result", result)
}

func dividedDiff(degree int, data []DataPoint, calculations *[][]float64, n *int) {
	if degree == 0 {
		for i := 0; i < *n-1; i++ {
			(*calculations)[0][i] = (data[i+1].Y - data[i].Y) / (data[i+1].X - data[i].X)
		}
		return
	}

	// this equation is used as it calculates the suitable no. of iterations for eac degree: {n - (degree + 1)}
	for i := 0; i < *n-(degree+1); i++ {
		(*calculations)[degree][i] = ((*calculations)[degree-1][i+1] - (*calculations)[degree-1][i]) / (data[i+degree].X - data[i].X)
	}
}

func calculate(calculations [][]float64, data []DataPoint, n *int) float64 {
	var result float64 = 0

	for i := 0; i < *n-1; i++ {
		if i == 0 {
			result += data[0].Y
			continue
		}

		// calculate values of Xs
		item := 1.0
		for j := 0; j < i; j++ {
			item *= x - data[j].X
		}
		// calculate values of Ys
		item *= calculations[i-1][0]

		// add it to the result
		result += item
	}

	return result
}

