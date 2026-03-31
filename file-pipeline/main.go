 // CodeCrafters — Operation Gopher Protocol
       // Module: File Pipeline
       // Author: Onminyi Andrew Okala
       // Squad:  The Interface


package main 

import (
	"fmt"
	"strings"
)

func main() {
	// reader := os.ReadFile("input.txt")
	// if err == nil {
	// 	fmt.Println("")
	// }

	// err := os.WriteFile("output.txt", )

	fmt.Println(capToTitle("Onminyi andrew okala"))
	fmt.Println(lowerToUpper("andrew okala onminyi"))
	fmt.Println(trimLeadTrail("Onminyi     Andrew    Okala"))
	fmt.Println(reverseStr("Reverse"))
}



func capToTitle(word string) string {
	return strings.Title(word)
	//words := strings.Fields(word)
	// for i, w := range words {
	// 	words[i] = strings.ToUpper(string(w[0])) + w[1:]
	// }
	// return  strings.Join(words, " ")
}

func lowerToUpper(word string) string {
	return strings.ToUpper(word)
}

func trimLeadTrail(word string) string {
	// words := strings.Fields(word)
	// wordss := strings.Join(words, " ")
	return strings.TrimSpace(word)
}

func reverseStr(word string) string {
	reverse := []rune(word)

	for i, j := 0, len(reverse) - 1; i < j; i, j = i + 1, j - 1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return string(reverse)
}