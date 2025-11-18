package main

import (
	"fmt"
)

func main() {
	b := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	c := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i := 0; i < len(b); i++ {
		c = append(c, "Пятница")
	}
	fmt.Println(c)
}
