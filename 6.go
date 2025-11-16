package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	if a > 0 && a%2 == 0 {
		fmt.Println("positive even")
	} else if a > 0 && a%2 != 0 {
		fmt.Println("positive odd")
	} else if a < 0 && a%2 == 0 {
		fmt.Println("neg even")
	} else {
		fmt.Println("neg odd")
	}
}
