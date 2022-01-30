package wordle_cli

import (
	"encoding/json"
	"log"
	"os"
)

type Words struct {
	W []string `json:"words"`
}

func loadWords(wsp *[]string) {
	b, err := os.ReadFile("words.json")
	if err != nil {
		log.Fatal(err)
	}

	var w Words
	if err = json.Unmarshal(b, &w); err != nil {
		log.Fatal(err)
	}
	*wsp = w.W
}
