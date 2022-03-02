package main

import (
	"fmt"
	"math"
)

func getFactorial(num int) int {
	if num > 1 {
		return num * getFactorial(num-1)
	} else {
		return 1
	}
}

func taylorPolinomial(x float64, order int) float64 {

	i := 1
	j := 3
	var max_val int = 20
	var result float64 = x

	for i <= order || order > max_val {

		if i%2 == 0 {
			result = result + (math.Pow(x, float64(j)) / float64(getFactorial(j)))
		} else {
			result = result - (math.Pow(x, float64(j)) / float64(getFactorial(j)))
		}
		i++
		j += 2
	}
	return result
}

func accurency(x float64, order int) float64 {
	if x > 0 {
		x = x - math.Ceil(x/(2*math.Pi))*(2*math.Pi)
	} else {
		x = x - math.Floor(x/(2*math.Pi))*(2*math.Pi)
	}

	if x >= 0 && x <= math.Pi/2 {
		x = x
	}
	if x >= math.Pi/2 && x <= math.Pi {
		x = math.Pi - x
	}
	if x >= math.Pi && x <= math.Pi*(3/2) {
		x = math.Pi - x
	}
	if x >= math.Pi*(3/2) && x <= math.Pi*(2) {
		x = x - math.Pi*(2)
	}

	return taylorPolinomial(x, order)
}

func main() {

	var x float64
	var order int

	fmt.Print("\nEnter the x of sin(x): ")
	_, err := fmt.Scanf("%f\n", &x)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("\nEnter the order: ")
	_, approximation := fmt.Scanf("%d\n", &order)

	if approximation != nil {
		fmt.Println(approximation)
	}

	fmt.Println("\nResult: ", accurency(x, order))
}
