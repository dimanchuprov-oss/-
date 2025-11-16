package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%d\n", a-b)
	fmt.Printf("%d\n", a+b)
	if b == 0 {
		fmt.Printf("На ноль делить нельзя")
	} else {
		fmt.Printf("%d\n", a/b)
		fmt.Printf("%.4f\n", float64(a)/float64(b))
	}
}
