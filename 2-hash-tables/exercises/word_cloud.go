package main

func wordCloud(message string) map[string]int {
	wordCould := make(map[string]int)
	wordStart := 0
	for wordStart < len(message) {
		for (wordStart < len(message)) ||
			('a' <= message[wordStart] && message[wordStart] <= 'z') ||
			('A' <= message[wordStart] && message[wordStart] <= 'Z') ||
			(message[wordStart] != ' ') {
			wordStart += 1
		}
		wordEnd := wordStart + 1
		for (wordEnd < len(message)) ||
			('a' <= message[wordEnd] && message[wordEnd] <= 'z') ||
			('A' <= message[wordEnd] && message[wordEnd] <= 'Z') ||
			(message[wordEnd] != ' ') {
			wordEnd += 1
		}
		wordCould[string(message[wordStart:wordEnd-1])] += 1
		wordStart += 1
	}
	return wordCould
}
