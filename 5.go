package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	a += 3
	fmt.Printf("%d\n", a)
	a += 2
	fmt.Printf("%d\n", a)
	a *= 4
	fmt.Printf("%d\n", a)
	a -= 10
	fmt.Printf("%d\n", a)
}
