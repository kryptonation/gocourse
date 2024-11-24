package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Investment Calculator")
	var investmentAmount float64 = 1000
	var roi = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1 + roi/100, years)
	fmt.Println(futureValue)
}