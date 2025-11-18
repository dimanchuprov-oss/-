package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	c := append([]byte{}, digitRegexp.Find(b)...)
	return c
}

func main() {
	fmt.Println(string(FindDigits("18_11_2025")))
	fmt.Println(string(CopyDigits("18_11_2025")))
}
