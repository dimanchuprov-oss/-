package main

import (
	"fmt"
)

func shellSort(a []int) {
	n := len(a)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := a[i]
			j := i
			for j >= gap && a[j-gap] > temp {
				a[j] = a[j-gap]
				j -= gap
			}
			a[j] = temp
		}
	}
}

func main() {
	var cnt int
	fmt.Scan(&cnt)
	arr := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			return
		}
	}

	shellSort(arr)

	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
