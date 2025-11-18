package main

import (
	"fmt"
)

func main() {
	b := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	c := b[1:]
	fmt.Println(c, len(b), cap(b), len(c))
}

