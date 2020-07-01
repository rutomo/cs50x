package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func Readability(text string) int {
	letters := 0
	words := 0
	sentences := 0

	for _, l := range text {
		if unicode.IsLetter(l) {
			letters++
		}
		if unicode.IsSpace(l) {
			words++
		}
		if strings.ContainsRune(".", l) || strings.ContainsRune("!", l) || strings.ContainsRune("?", l) {
			sentences++
		}
	}

	L := float64(letters) / float64(words) * 100
	S := float64(sentences) / float64(words) * 100

	index := math.Round(0.0588*L - 0.296*S - 15.8)

	return int(index)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text: ")
	text, _ := reader.ReadString('\n')

	index := Readability(text)

	if index < 1 {
		fmt.Printf("Before Grade 1")
	} else if index >= 16 {
		fmt.Printf("Grade 16+")
	} else {
		fmt.Printf("Grade %d", int(index))
	}
}
