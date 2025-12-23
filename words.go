package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	text, _ := reader.ReadString('\n')
	lower := strings.ToLower(text)
	delimiters := " .,!?;:()[]{}<>\"'`-–—\t\n\r"
	for _, delim := range delimiters {
		text = strings.ReplaceAll(lower, string(delim), " ")
	}
	
	words := strings.Fields(lower)
	
	uniqueWords := make(map[string]bool)
	for _, word := range words {
		uniqueWords[word] = true
	}
	var result []string
	for word := range uniqueWords {
		result = append(result, word)
	}

	fmt.Println("Уникальные слова:", result)
}
