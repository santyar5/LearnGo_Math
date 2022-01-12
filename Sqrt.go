package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	lim := 1.0
	for i := 0; i < 10; i++ {
		lim -= (lim*lim - x) / (2 * lim)
	}
	return lim
}

func Pow(x float64) float64 {
	return x * x
}

func main() {
	fmt.Println("Введите число:")
	var num int
	var _, err = fmt.Scan(&num)
	if err != nil {
		fmt.Printf("Ошибка ввода: %s", err)
	} else {
		fmt.Printf("Pow & Sqrt num %d - result %f", num, Pow(Sqrt(float64(num))))
	}
}
