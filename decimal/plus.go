package decimal

// Plus two positive integers
func Plus(a string, b string) string {
	var carry byte
	var res []byte
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 && j >= 0 {
		sum := a[i] - '0' + b[j] - '0' + carry
		res = append(res, sum%10+'0')
		carry = sum / 10
		i--
		j--
	}

	for i >= 0 {
		sum := a[i] - '0' + carry
		res = append(res, sum%10+'0')
		carry = sum / 10
		i--
	}

	for j >= 0 {
		sum := b[j] - '0' + carry
		res = append(res, sum%10+'0')
		carry = sum / 10
		j--
	}

	if carry > 0 {
		res = append(res, carry+'0')
	}

	for left, right := 0, len(res)-1; left < right; left, right = left+1, right-1 {
		res[left], res[right] = res[right], res[left]
	}
	return string(res)
}
