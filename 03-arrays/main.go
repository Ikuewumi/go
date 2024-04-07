package main

import (
	"fmt"
	"math"
)

func main() {
	x := []int{98, 93, 77, 82, 83 }
	

    y := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}


	fmt.Println(
		getMininumInteger(x),
        getAverage(y),
    )

}

func getMininumInteger(array []int) int {
	var min int = math.MaxInt

	for _, value := range array {
		if value < min {
			min = value
		}
	}

	return min

}

func getAverage(array []int) float64 {
	var total float64

	for _, value := range array {
		total += float64(value)
	}

	average := total / float64(len(array))

	return average
}
