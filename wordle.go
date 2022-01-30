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
	sc.Split(bufio.ScanWords)

	loadWords(&words)

	alphabet := map[rune]int{}
	var keys []rune
	c := 'a'
	for i := 0; i < 26; i++ {
		tmp := rune(int(c) + i)
		keys = append(keys, tmp)
		alphabet[tmp] = UNCHECKED
	}

	//rand.Seed(time.Now().UnixNano())
	wordle := words[rand.Intn(len(words))]
	fmt.Println("wordle:", wordle)

	for i := 0; i < 6; i++ {
		ans := InputWord()

		res := CompareWordle(wordle, ans)

		fmt.Printf("%d: ", i+1)
		for j, v := range res {
			alphabet[rune(ans[j])] = v
			switch v {
			case UNUSED:
				fmt.Printf("\x1b[41m%c\x1b[0m ", ans[j])
			case BITE:
				fmt.Printf("\x1b[43m%c\x1b[0m ", ans[j])
			case EAT:
				fmt.Printf("\x1b[42m%c\x1b[0m ", ans[j])
			}
		}
		j := 0
		for _, v := range keys {
			if j%10 == 0 {
				fmt.Println()
			}
			switch alphabet[v] {
			case UNUSED:
				fmt.Printf("\x1b[41m")
			case BITE:
				fmt.Printf("\x1b[43m")
			case EAT:
				fmt.Printf("\x1b[42m")
			}
			fmt.Printf("%c\u001B[0m ", v)
			j++
		}
		fmt.Printf("\n\n")

		if ans == wordle {
			fmt.Println("correct!!")
			break
		}
	}
}
