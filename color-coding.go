package wordle_cli

import "fmt"

func printCharWithStatus(status int, c rune) {
	switch status {
	case UNUSED:
		fmt.Printf("\x1b[41m%c\x1b[0m ", c)
	case BITE:
		fmt.Printf("\x1b[43m%c\x1b[0m ", c)
	case EAT:
		fmt.Printf("\x1b[42m%c\x1b[0m ", c)
	}
}

func printAlphabet(alphabet map[rune]int) {
	for i := 0; i < 26; i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		switch alphabet[rune('a'+i)] {
		case UNUSED:
			fmt.Printf("\x1b[41m")
		case BITE:
			fmt.Printf("\x1b[43m")
		case EAT:
			fmt.Printf("\x1b[42m")
		}
		fmt.Printf("%c\u001B[0m ", rune('a'+i))
	}
}

func printHistory(history [][5]int) {
	for _, v := range history {
		for _, vv := range v {
			switch vv {
			case UNUSED:
				fmt.Printf("\x1b[41m  \x1b[0m ")
			case BITE:
				fmt.Printf("\x1b[43m  \x1b[0m ")
			case EAT:
				fmt.Printf("\x1b[42m  \x1b[0m ")
			}
		}
		fmt.Println()
	}
}
