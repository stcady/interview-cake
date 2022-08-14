package main

func hasPalindromePermutation(word string) bool {
	unpairedCharacters := make(map[rune]bool)
	for _, character := range word {
		if _, ok := unpairedCharacters[character]; ok {
			delete(unpairedCharacters, character)
		} else {
			unpairedCharacters[character] = true
		}
	}
	return len(unpairedCharacters) < 2
}
