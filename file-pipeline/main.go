/*
// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [The Interface]
// ───────────────────────────────────────────
// Input line types:
//   [
        Lines in ALL CAPS
        Lines in all lowercase
        Lines with extra leading/trailing spaces
        Lines that are only dashes or blanks
       ]
//
// Transformation rules (in order):
//   1. [Convert ALL CAPS lines to Title Case]
//   2. [Convert all lowercase lines to uppercase]
//   3. [Trim all leading and trailing whitespace]
//   4. [Reverse the words in any line that contains the word REVERSE]
//   5. [Remove lines that are only dashes or blanks]
//
// Output format:
//   [Header: yes]
//   [Line numbering format: "1"]
//   [Summary block: yes]
//
// Terminal summary fields:
//   [Lines read    : [number]]
//   [Lines written : [number]]
//   [Lines removed : [number]]
//   [Rules applied : [list your 5 rules]]
// ═══════════════════════════════════════════
═════════════════════════════════════════
*/

package main

import (
	"os"
	"strings"
)

func capToTitle(word string) string {
	if word != "" {
		words := strings.Fields(word)
	for i, w := range words {
		words[i] = strings.ToUpper(string(w[0])) + w[1:]
	}
	return strings.Join(words, " ")
	}
	return  word
}

func lowerToUpper(word string) string {
	if word != "" {
		return strings.ToUpper(word)
	}
	return word
}

func trimLeadTrail(word string) string {
	words := strings.Fields(word)
	wordN := strings.Join(words, " ")
	return strings.TrimSpace(wordN)
}

func reverseStr(word string) string {

	if strings.Contains(word, "REVERSE") {
		words := strings.Fields(word)
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
	}
	return word
}

func truncate(word string) string {
	if len(word) > 80 {
		return word + "[TRUNCATED]"
	}
	return word
}

func receiver(word string) string {
	word = capToTitle(word)
	word = lowerToUpper(word)
	word = trimLeadTrail(word)
	word = reverseStr(word)
	word = truncate(word)

	return word
}

func main() {
	in, _ := os.ReadFile("input.txt")
	word := strings.Split(strings.ReplaceAll(string(in), "\r\n", "\n"), "\n")
	for i := range word {
		word[i] = receiver(word[i])
	}
	os.WriteFile("output.txt", []byte(strings.Join(word, "\n")), 0644)
}