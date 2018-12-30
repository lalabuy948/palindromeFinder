package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func longest(s []string) string {
	longest := ""
	for _, word := range s {
		if len(longest) <= len(word) {
			longest = word
		}
	}
	return longest
}

func main() {
	file, err := os.Open("words_rus.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var palindrome []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == reverse(scanner.Text()) {
			palindrome = append(palindrome, scanner.Text())
		}
	}

	fmt.Println(longest(palindrome))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
