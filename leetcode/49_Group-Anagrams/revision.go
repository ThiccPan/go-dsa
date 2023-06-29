package main

type Key [26]byte

func strKey(str string) (key Key) {
	for v := range str {
		key[str[v]-'a']++
	}
	return key
}

func GroupAnagramsRevision(strs []string) [][]string {
	subStrMap := make(map[Key][]string)
	subStrs := [][]string{}

	for _, str := range strs {
		key := strKey(str)
		subStrMap[key] = append(subStrMap[key], str)
	}

	for _, v := range subStrMap {
		subStrs = append(subStrs, v)
	}
	return subStrs
}