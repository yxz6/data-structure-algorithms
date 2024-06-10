package string

import "testing"

func TestKMP(t *testing.T) {
	t.Log(KMP("ababababcab", "ababcab"))
	t.Log(KMP("abababeabcab", "ababcab"))
}
