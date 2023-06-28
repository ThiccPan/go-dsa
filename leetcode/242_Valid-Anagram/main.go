package main

import "fmt"

// my first try
func isAnagram(s string, t string) bool {
	strMap := map[rune]int{}

	if len(s) != len(t) {
		return false
	}

	// map every rune in str
	for idx, strRune := range s {
		strMap[strRune]++
		strMap[rune(t[idx])]--
	}

	// check if any rune is occuring odd times
	// if so then not anagram

	// unoptimized checking
	// for _, v := range strMap {
	// 	if v % 2 != 0 || v < 0{
	// 		return false
	// 	}
	// }

	for _, v := range strMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	first := isAnagram("hello", "eollh")
	fmt.Println(first)
}
