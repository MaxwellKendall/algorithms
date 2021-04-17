package exercise0103

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

