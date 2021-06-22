package recursion

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var res []byte
	temp, carry, length := 0, 0, len(num1)+len(num2)-2
	for i := len(num2) - 1; i >= 0; i-- {
		for j := len(num1) - 1; j >= 0; j-- {
			temp = int(num2[i]-'0')*int(num1[j]-'0') + carry
			if len(res) > length-i-j {
				temp += int(res[length-i-j] - '0')
				carry = temp / 10
				res[length-i-j] = byte(temp%10 + int('0'))
			} else {
				carry = temp / 10
				res = append(res, byte(temp%10+int('0')))
			}
		}
		if carry != 0 {
			res = append(res, byte(carry+int('0')))
		}
		carry = 0
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
