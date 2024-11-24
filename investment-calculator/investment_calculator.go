package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Investment Calculator")
	const inflationRate = 3.5
	var investmentAmount float64
	var roi float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter ROI: ")
	fmt.Scan(&roi)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + roi/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}