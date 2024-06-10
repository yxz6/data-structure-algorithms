package string

// KMP time complexity O(n+m), space complexity O(m)
func KMP(s, pattern string) int {
	// next[0] = 0
	// [0, next[i]] = [i-next[i], i]
	next := make([]int, len(pattern))
	next[0] = -1
	for i := 1; i < len(next); i++ {
		j := i - 1
		for next[j] >= 0 {
			if pattern[i-1] == pattern[next[j]] {
				next[i] = next[j] + 1
				break
			} else {
				j = next[j]
			}
		}
	}

	i := 0
	j := 0
	for j < len(s) {
		if s[j] == pattern[i] {
			i++
			j++
			if i == len(pattern) {
				return j - len(pattern) + 1
			}
		} else {
			i = next[i]
			if i == -1 {
				i = 0
				j++
			}
		}
	}
	return -1
}
