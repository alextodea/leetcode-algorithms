package strings

type TargetNames struct {
	name  string
	typed string
}

func (n *TargetNames) isLongPressed() bool {
	lengthOfName := len(n.name)
	lengthOfTyped := len(n.typed)
	j := 0

	for i := 0; i < lengthOfTyped; i++ {
		if j < lengthOfName && (n.name[j] == n.typed[i]) {
			j++
		} else if i == 0 || (n.typed[i] != n.typed[i-1]) {
			return false
		}
	}

	return j == lengthOfName
}
