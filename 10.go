package main

import "fmt"

func divmod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("на ноль делить нельзя!")
	}
	return a / b, a % b, nil
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	x, y, err := divmod(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x, y)
	}
}
