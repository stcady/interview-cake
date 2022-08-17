package main

func getClosingParen(sentence string, openPosition int) int {
	nestedParen := 0
	subSentence := sentence[openPosition+1:]
	for i, char := range subSentence {
		if char == '(' {
			nestedParen += 1
		} else if char == ')' {
			if nestedParen == 0 {
				return i
			} else {
				nestedParen -= 1
			}
		}
	}
	return 0
}
