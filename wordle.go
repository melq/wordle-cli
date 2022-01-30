package wordle_cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

const (
	UNCHECKED = iota
	UNUSED
	BITE
	EAT
)

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

func PlayWordle() {
	sc.Split(bufio.ScanWords)

	var words []string
	loadWords(&words)

	alphabet := map[rune]int{}
	c := 'a'
	for i := 0; i < 26; i++ {
		alphabet[c] = 0
	}

	rand.Seed(time.Now().UnixNano())
	wordle := words[rand.Intn(len(words))]
	fmt.Println(wordle)

	for i := 0; i < 6; i++ {
		ans := next()
		if ans == wordle {
			fmt.Println("correct!!")
			break
		} else {
			res := CompareWordle(wordle, ans)
			for i, _ := range res {

			}
		}
	}
}
