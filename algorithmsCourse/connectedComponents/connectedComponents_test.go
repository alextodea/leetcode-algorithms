package connectedComponents

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	top5CC := main()
	fmt.Print(top5CC)
	if 1 != 2 {
		t.Error("sss")
	}
}
