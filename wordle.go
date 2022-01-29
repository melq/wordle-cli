package wordle_cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

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

	answer := words[rand.Intn(len(words))]
	fmt.Println(answer)
	for i := 0; i < 6; i++ {
		prop := next()
		if prop == answer {
			fmt.Println("correct!!")
			break
		}
	}
}
