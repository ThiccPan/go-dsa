package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	strStore := [][]string{}
	strMap := map[string][]string{}
	for _, str := range strs {
		strMap[sortStr(str)] = append(strMap[sortStr(str)], str)
	}
	fmt.Println(strMap)
	for _, strs := range strMap {
		strStore = append(strStore, strs)
	}
	return strStore
}

func sortStr(str string) string {
	splitStr := strings.Split(str, "")
	sort.Strings(splitStr)
	sorted := strings.Join(splitStr, "")
	return sorted
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
}