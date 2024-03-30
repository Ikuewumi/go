package main

import (
	"fmt"
	"os"
)


func main () {
    fizzBuzzProblem()
    fmt.Println("")
    os.Exit(0)
}



func fizzBuzzProblem () {
    var output string
    
    for i := 0; i <= 100; i++ {
        output = ""

        if i % 3 == 0 && i % 5 == 0  {
            output = "FizzBuzz"
        } else if i % 3 == 0 {
            output = "Fizz"
        } else if i % 5 == 0 {
            output = "Buzz"
        }

        if (len(output) > 0) {
            fmt.Printf("%s \n", output)
        } else {
            fmt.Printf("%d \n", i)
        }

    }
}


func printNumbersDivisibleByThree () {
    for i := 0; i < 100; i++ {
        if (i % 3 == 0) {
            fmt.Printf("%d, ", i)
        }
    }
}
