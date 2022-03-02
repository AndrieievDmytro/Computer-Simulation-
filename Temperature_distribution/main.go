package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Print matrtix
func MatPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func main() {
	calculateTempreture()
}

func calculateTempreture() {
	m := 40
	n := 40
	A := mat.NewDense(m*n, m*n, nil)
	var equation_matrix [][]float64
	var vector []float64
	values := []float64{150, 50.0, 200.0, 100.0}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			equation := make([]float64, m*n)
			edge := 0.0
			equation[i*n+j] = -4
			// bottom boundary
			if i+1 == m {
				edge += -values[2]
			} else {
				equation[(i+1)*n+j] = 1
			}
			// top boundary
			if i-1 < 0 {
				edge += -values[0]
			} else {
				equation[(i-1)*n+j] = 1
			}
			// right boundary
			if j+1 == n {
				edge += -values[1]
			} else {
				equation[i*n+j+1] = 1
			}
			//  left boundary
			if j-1 < 0 {
				edge += -values[3]
			} else {
				equation[i*n+j-1] = 1
			}
			// fmt.Println(equation, " ", edge)
			equation_matrix = append(equation_matrix, equation)
			vector = append(vector, edge)
		}
	}
	for idx, eq := range equation_matrix {
		A.SetRow(idx, eq)
		// fmt.Println(equation_matrix, " ", vector)
		// fmt.Println(eq, vector[idx])
	}
	v := mat.NewVecDense(n*m, vector)
	err := A.Inverse(A)
	if err != nil {
		fmt.Println("ERROR: not inverteble matrix")
	}
	v.MulVec(A, v)
	// matPrint(v)

	fmt.Print(",")
	for j := 0; j < n; j++ {
		fmt.Print(values[2], ",")
	}
	for i := m; i > 0; i-- {
		fmt.Println()
		fmt.Print(values[3], ",")
		for j := 0; j < n; j++ {
			fmt.Print(v.At((i-1)*n+j, 0), ",")
		}
		fmt.Print(values[1])
	}
	fmt.Println()
	fmt.Print(",")
	for j := 0; j < n; j++ {
		fmt.Print(values[0], ",")
	}
	fmt.Println()
}
