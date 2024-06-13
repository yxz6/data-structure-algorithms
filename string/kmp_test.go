package string

import "testing"

func TestKMP(t *testing.T) {
	t.Log(KMP("ababababcab", "aabaaccaabaabaabd"))
	t.Log(KMP("abababeabcab", "ababcab"))
}
