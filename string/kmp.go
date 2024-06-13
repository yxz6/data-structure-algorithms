package string

import "fmt"

// KMP time complexity O(n+m), space complexity O(m)
func KMP(s, pattern string) int {
	lps := createLongestPrefixTable(pattern)
	fmt.Println(lps)
	//// next[0] = 0
	//// [0, next[i]-1] = [i-next[i], i-1]
	//next := make([]int, len(pattern))
	//for i := 1; i < len(next); i++ {
	//	j := i - 1
	//	for next[j] >= 0 {
	//		if pattern[i-1] == pattern[next[j]] {
	//			next[i] = next[j] + 1
	//			break
	//		} else {
	//			j = next[j]
	//		}
	//	}
	//}
	//
	//i := 0
	//j := 0
	//for j < len(s) {
	//	if s[j] == pattern[i] {
	//		i++
	//		j++
	//		if i == len(pattern) {
	//			return j - len(pattern) + 1
	//		}
	//	} else {
	//		i = next[i]
	//		if i == -1 {
	//			i = 0
	//			j++
	//		}
	//	}
	//}
	return -1
}

func createLongestPrefixTable(pattern string) []int {
	// example:
	/*
	   	a a b a a c c a a b a a b a a b d
	   	0 1 0 1 2 0 0 1 2 3 4 5 3 1 2 3 0
	                            if b==c, then prefix len 6=5+1
	                            as b!=c, lps[5-1]=2, pattern[lps[5-1]]=b, so lps=lps[5-1]+1

	*/

	lps := make([]int, len(pattern))
	for i := 1; i < len(pattern); i++ {
		prevPrefixLen := lps[i-1]
	LOOP:
		for {
			// prevPrefix is pattern[0, prevPrefixLen-1], pattern[prevPrefixLen] is the next char
			// pattern[0, prevPrefixLen-1] == pattern[i-prevPrefixLen, i-1]
			if pattern[prevPrefixLen] == pattern[i] {
				// pattern[0, prevPrefixLen] == pattern[i-prevPrefixLen, i]
				// so prefixLen = prevPrefixLen+1
				lps[i] = prevPrefixLen + 1
				break LOOP
			} else {
				if prevPrefixLen == 0 {
					break LOOP
				}
				// try to find longest prefix in sub string
				prevPrefixLen = lps[prevPrefixLen-1]
			}
		}
	}
	return lps
}
