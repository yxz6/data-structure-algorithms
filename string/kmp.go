package string

// KMP time complexity O(n+m), space complexity O(m)
func KMP(s, pattern string) int {
	lps := createLongestPrefixTable(pattern)
	i, j := 0, 0
	for i < len(s) && j < len(pattern) {
		if pattern[j] == s[i] {
			i++
			j++
		} else if j > 0 {
			//previous string's longest prefix, and go to the next char
			j = lps[j-1]
		} else {
			i++
		}
	}
	if j == len(pattern) {
		return i - j
	}
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
				// try to find the longest prefix in sub string
				prevPrefixLen = lps[prevPrefixLen-1]
			}
		}
	}
	return lps
}
