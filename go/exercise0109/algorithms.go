// String Rotation: Assume you have a method isSubstringwhich checks if oneword
// is a substring of another. Given two strings, sl and s2, write code to check
// if s2 is a rotation of sl using only one call to isSubstring (e.g.,
// "waterbottle" is a rotation of "erbottlewat".
package exercise0109

import (
	"strings"
)

// Time: O(n),
//       O(1) in calls to isSubstring.
// Space: O(n)
func Double(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return isSubstring(s2, s1+s1)
}

// Time: O(n^2)
// Space: O(1)
func Loop(s1, s2 string) bool {
	for i, _ := range s1 {
		if strings.HasPrefix(s2, s1[i:]) {
			return strings.HasSuffix(s2, s1[:i])
		}
	}
	return false
}
