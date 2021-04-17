# URLify Algorithm Analysis

> Given a string, replace all spaces with %20

## Solution:

```go
func ReplaceSpace(str string) string {
	var rtrn string
	for _, v := range str {
		if string(v) == string(" ") {
			rtrn += string("%20")
		} else {
			rtrn += string(v)
		}
	}
	return rtrn
}
```

## Time & Time Complexity: O(n)
Both space and time complexity increase as the input 
size increases.

## Conclusion
A sorted string would help in this case. We could go down to, at worst, O(log n) assuming we know the sorting algorithm. 