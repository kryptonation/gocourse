package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Investment Calculator")
	const inflationRate = 3.5
	var investmentAmount float64 = 1000
	roi := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow(1 + roi/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}