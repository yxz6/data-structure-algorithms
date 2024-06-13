package string

import "testing"

func TestKMP(t *testing.T) {
	t.Log(KMP("acaabaaccaabaabaabde", "aabaaccaabaabaabd"))
	t.Log(KMP("abababeabcab", "ababcab"))
	t.Log(KMP("abababeabcab", "abc"))
}
