package wordle_cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

const (
	UNCHECKED = iota
	UNUSED
	BITE
	EAT
)

var words []string

type Words struct {
	W []string `json:"words"`
}

func loadWords(w *[]string) {
	b, err := os.ReadFile("words.json")
	if err != nil {
		log.Fatal(err)
	}

	var words Words
	if err = json.Unmarshal(b, &words); err != nil {
		log.Fatal(err)
	}
	*w = words.W
}

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

func InputWord() string {
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

func PlayWordle() {
	var history [][5]int
	alphabet := map[rune]int{}

	sc.Split(bufio.ScanWords)

	loadWords(&words)

	c := 'a'
	for i := 0; i < 26; i++ {
		tmp := rune(int(c) + i)
		alphabet[tmp] = UNCHECKED
	}

	//rand.Seed(time.Now().UnixNano())
	wordle := words[rand.Intn(len(words))]
	fmt.Println("wordle:", wordle)

	for i := 0; i < 6; i++ { // try six times
		ans := InputWord()
		res := evaluateAnswer(wordle, ans)
		history = append(history, res)

		fmt.Printf("%d: ", i+1)
		for j, v := range res { // print answer color
			alphabet[rune(ans[j])] = v
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
