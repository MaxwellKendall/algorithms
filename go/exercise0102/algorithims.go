// CheckPermutation: Implement an algorithm to determine if one string is a permutation of another.
// Requirements: Notate space and time complexity

// Time complexity:
// 1. Can we get less than 0(n)? --> No.
// 2. Will sorting help? --> No.

// Space complexity:
// 1. Can we get less than 0(n)? --> Yes.
// 2. Will sorting help? --> Yes.

package exercise0102

// WithMap complexity estimates:
// Time: O(n)
// Space: O(1) we are at 0(1) because of the finite nature of the alphabet.
func WithMap(s1 string, s2 string) bool {
	if (len(s1) != len(s2)) {
		return false
	}
	myMap := map[rune]int{}
	for _, v := range s1 {
		myMap[v]++
	}
	for _, v := range s2 {
		myMap[v]--
	}
	for _, v := range myMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func buildDynamicMap(str string) map[rune]int {
	rtrn := map[rune]int{}
	for _, v := range str {
		rtrn[v]++
	}
	return rtrn
}

// With2Maps complexity estimates:
// Time: O(n)
// Space: O(1) we are at 0(1) because of the finite nature of the alphabet.
func With2Maps(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	firstMap := buildDynamicMap(s1)
	secondMap := buildDynamicMap(s2)

	for key, v := range firstMap {
		if secondMap[key] != v {
			return false
		}
	}
	return true
}
