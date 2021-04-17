# Check Permutation Algorithm Analysis

Function to answer the question:

> Given two strings, can they be arranged in such a way that they are the same word. In other words, is one string a "permutation" of the other.

## Solution #1: WithMap

```go
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
```

### Time Complexity Estimate: O(n)
Here we iterate over each character in each string once. This means the time complexity is coupled to the size of the input.
Not bad.

### Space Complexity Estimate: O(1)
No matter how big the input size, we are dealing with a finite data set. When
we have repeat characters, the memory used to capture these characters is the same
as the first time we stored the character. The `1` here in this case, then, refers 
to the number of items in a given alphabet. In our case, the english alphabet
which is 26 characters long. The memory we need for this algorithm, then, is 
the same as the size of each unique character in the alphabet.

## Solution #2: With2Maps

```go
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
```

### Time Complexity Estimate: O(n)
Same as above, although optimal solution for time as we only do one iteration rather than two.
### Space Complexity Estimate: O(1)
Same as above, although sub optimal as two maps are stored rather than one.

## Conclusion: WithMaps for Space, With2Maps for Time