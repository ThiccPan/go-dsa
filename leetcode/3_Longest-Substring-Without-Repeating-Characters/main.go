package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	left := 0
	charlist := map[string]int{}

	for right := 0; right < len(s); right++ {
		rightVal := string(s[right])
		_, ok := charlist[rightVal];

		// case 1 : char is unique, then add to map
		if !ok || charlist[rightVal] < left{
			charlist[rightVal] = right

			// default : define max by subtract left with i
			if right-left+1 > max {
				max = right - left + 1
			}
			fmt.Println("case 1", left, right)
			continue
		}

		// case 2 : char not unique, move start to the right until cur substr is unique
		if _, ok := charlist[rightVal]; ok {
			left = charlist[rightVal] + 1
			fmt.Println("case 2", left, right)
			charlist[rightVal] = right
			// default : define max by subtract left with i
			if right-left+1 > max {
				max = right - left + 1
			}
		}

	}
	// fmt.Println(charlist)
	return max
}
