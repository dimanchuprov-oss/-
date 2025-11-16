package main

import "fmt"

func main() {
	var a [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}
	sum := 0
	ind := 0
	k := 0
	for i := 0; i < 5; i++ {
		sum += a[i]
		if a[i] > k {
			ind = i
			k = a[i]
		}
	}
	fmt.Println(sum)
	fmt.Println(ind, k)
	var result []int
	for _, num := range a {
		if num > 0 {
			result = append(result, num)
		}
	}
	if len(result) == 0 {
		fmt.Println("")
	} else {
		for i, num := range result {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	}
}
