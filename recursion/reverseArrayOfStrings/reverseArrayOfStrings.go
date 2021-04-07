package recursion

func reverseString(s []byte) {
	if len(s) == 1 {
		return
	}

	reverseRecursively(0, s)
}

var temp byte

func reverseRecursively(i int, s []byte) {
	if i >= len(s)-i {
		return
	}

	temp = s[i]
	s[i] = s[len(s)-1-i]
	s[len(s)-1-i] = temp
	i++

	reverseRecursively(i, s)
}
