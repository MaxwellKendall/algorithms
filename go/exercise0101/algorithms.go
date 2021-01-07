// Is Unique: Implement an algorithm to determine if a string has all unique
// characters. What if you cannot use additional data structures?
//
// Observations:
// - Cannot outperform O(n) time since we have to look at every character.
package exercise0101

import (
	"sort"
)

// Time: O(n)
// Space: O(n)
// Uses additional data structure.
func WithMap(s string) bool {
	record := map[rune]bool{}
	for _, c := range s {
		if _, ok := record[c]; ok {
			return false
		}
		record[c] = true
	}
	return true
}

// Time: O(n^2)
// Space: O(1)
// Does not use additional data structure.
func NestedLoop(s string) bool {
	for i, c1 := range s {
		for _, c2 := range s[i+1:] {
			if c1 == c2 {
				return false
			}
		}
	}
	return true
}

// Time: O(n log n)
// Space: O(1)
// Does not use an additional data structure.
// Destructive.
func sortFirst(s []rune) bool {
	if len(s) < 1 {
		return true
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	for i, _ := range s[1:] {
		if s[i] == s[i+1] {
			return false
		}
	}
	return true
}

// This only makes sense with a mutable data structure, and strings are
// not mutable in go, so this function wraps our actual function in a
// consistent interface.
func SortFirst(s string) bool {
	return sortFirst([]rune(s))
}
