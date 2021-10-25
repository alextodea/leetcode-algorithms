package strings

import "testing"

func Test1(t *testing.T) {
	input := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	result := numUniqueEmails(input)
	if result != 2 {
		t.Error("result should be 2")
	}
}

func Test2(t *testing.T) {
	input := []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	result := numUniqueEmails(input)
	if result != 3 {
		t.Error("result should be 3")
	}
}
