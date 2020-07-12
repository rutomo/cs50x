package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func is_integer(key string) bool {

	if _, err := strconv.Atoi(key); err == nil {
		return true
	} else {
		return false
	}
}

func main() {

	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Print("Usage: ./caesar key")
		return
	}
	if !is_integer(os.Args[1]) {
		fmt.Print("Usage: ./caesar key")
		return
	}

	key, err := strconv.Atoi(os.Args[1])
	var ciphertext string = ""
	if err == nil {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("plaintext: ")
		text, _ := reader.ReadString('\n')

		for _, l := range text {
			ciphertext += string(int(l) + (key % 26))
		}
		fmt.Printf("ciphertext: %s", ciphertext)
	}

	return
}
