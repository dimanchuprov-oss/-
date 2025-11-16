package main

import "fmt"

func main() {
	var a int
	var b int = 0
	fmt.Scanf("%d", &a)

	for i := 0; i <= a; i++ {
		b += i
	}
	fmt.Printf("%d\n", b)
}
