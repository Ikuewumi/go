package main

import (
	"fmt"
	"os"
)

func main() {
	getCelsiusFromFarenheit()
	getMetersFromFeet()
    os.Exit(0)
}


func getMetersFromFeet() {
    fmt.Print("Enter the Distance in feet: ")
    var feet float64
    fmt.Scanf("%f", &feet)

    meters := feet * 0.3048

    fmt.Printf("%0.1fft has been converted to %0.1fmeters \n", feet, meters)
}

func getCelsiusFromFarenheit() {
	fmt.Print("Enter the Temperature in Farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9

	fmt.Printf("%0.1fF has been converted to %0.1fC \n", input, output)

}
