package main

import (
	"fmt"
)

func main() {
	b := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(b)
	fmt.Println(b[1])
	fmt.Println(b[2:])
}
