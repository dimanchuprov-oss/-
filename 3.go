package main

import "fmt"

func main() {
	const pi float64 = 3.14159
	var r float64
	fmt.Scanf("%f\n", &r)
	area := pi * r * r
	fmt.Printf("%.3f\n", area)
}
