package wordle_cli

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func isInWords(word string) bool {
	for _, v := range words {
		if word == v {
			return true
		}
	}
	return false
}

func submitWord() string {
	for {
		fmt.Println("Submit a five-letter word:")
		word := next()
		if len(word) != 5 {
			continue
		} else if isInWords(word) {
			return word
		} else {
			fmt.Println("Not in word list")
		}
	}
}
