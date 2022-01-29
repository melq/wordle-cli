package wordle_cli

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
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
	var words []string
	loadWords(&words)

	answer := words[rand.Intn(len(words))]
	fmt.Println(answer)
}
