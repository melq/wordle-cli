package wordle_cli

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	UNCHECKED = iota
	UNUSED
	BITE
	EAT
)

var words []string

func PlayWordle() {
	var history [][5]int
	alphabet := map[rune]int{}

	loadWords(&words)

	c := 'a'
	for i := 0; i < 26; i++ {
		tmp := rune(int(c) + i)
		alphabet[tmp] = UNCHECKED
	}

	rand.Seed(time.Now().UnixNano())
	wordle := words[rand.Intn(len(words))]

	for i := 0; i < 6; i++ {
		ans := submitWord()
		res := evaluateAnswer(wordle, ans)
		history = append(history, res)

		fmt.Printf("%d: ", i+1)
		for j, v := range res {
			if alphabet[rune(ans[j])] != EAT {
				alphabet[rune(ans[j])] = v
			}
			printCharWithStatus(v, rune(ans[j]))
		}

		printAlphabet(alphabet)
		fmt.Printf("\n\n")

		if ans == wordle {
			fmt.Println("correct!")
			printHistory(history)
			break
		}
	}
}
