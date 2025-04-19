package models

import "fmt"

func Somar(a, b float64) float64 {
	return a + b
}

func Subtracao(a, b float64) float64 {
	return a - b
}

func Multiplicacao(a, b float64) float64 {
	return a * b
}

func Divisão(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}
