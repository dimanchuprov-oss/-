package main

import "fmt"

func division(a int, b int) (float64, bool) {
	if b == 0 {
		return 0, false
	} else {
		c := float64(a) / float64(b)
		return c, true
	}
}

func main() {
	var a, b int

	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)
	result, ok := division(a, b)
	if ok {
		fmt.Printf("Результат: %.2f\n", result)
	} else {
		fmt.Printf("Деление на ноль(")
	}
}