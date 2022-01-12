package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	lim := 2.0
	var result float64
	for i := 0; i < 10; i++ {
		var res = (lim + (x / lim)) / 2
		//fmt.Println("res = ", res)
		//fmt.Println(lim/res)
		if lim/res > 1 {
			lim = res
		} else {
			result = res
			return result
		}
		result = res
	}
	return result
}

func Pow(x float64) float64 {
	return x * x
}

func main() {
	fmt.Println("Введите число:")
	var num int
	var _, err = fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	} else {
		fmt.Println(Pow(Sqrt(float64(num))))
	}
}
