package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	swap(&a, &b)

	fmt.Println(a, b)
}
