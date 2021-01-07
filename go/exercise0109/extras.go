package exercise0109

import "strings"

func isSubstring(s1, s2 string) bool {
	return strings.Contains(s2, s1)
}
