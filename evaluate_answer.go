package wordle_cli

func contains(str string, c rune) bool {
	for _, v := range str {
		if c == v {
			return true
		}
	}
	return false
}

func evaluateAnswer(wordle string, ans string) [5]int {
	res := [5]int{}
	for i, v := range ans {
		if v == rune(wordle[i]) {
			res[i] = EAT
		} else if contains(wordle, v) {
			res[i] = BITE
		} else {
			res[i] = UNUSED
		}
	}

	return res
}
