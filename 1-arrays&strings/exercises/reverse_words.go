package main

func reverseWords(message []rune) []rune {
	message = reverseCharacters(message, 0, len(message)-1)
	currentWordStart := 0
	i := 0
	for i < len(message)+1 {
		if (i == len(message)) || (message[i] == ' ') {
			message = reverseCharacters(message, currentWordStart, i-1)
			currentWordStart = i + 1
		}
		i += 1
	}
	return message
}

func reverseCharacters(message []rune, left int, right int) []rune {
	for left < right {
		message[left], message[right] = message[right], message[left]
		left += 1
		right -= 1
	}
	return message
}
