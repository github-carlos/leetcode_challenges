package main

import (
	"fmt"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	for _, c := range magazine {
		if strings.Contains(ransomNote, string(c)) {
			ransomNote = strings.Replace(ransomNote, string(c), "", 1)
		}
	}
	return len(ransomNote) == 0
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "aab"))
}