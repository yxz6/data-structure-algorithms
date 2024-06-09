package decimal

// Minus two positive integers
func Minus(a string, b string) string {
	if len(a) < len(b) || (len(a) == len(b) && a < b) {
		delta := Minus(b, a)
		if delta[0] == '-' {
			return delta[1:]
		}
		return "-" + delta
	}

	var borrow int8
	var res []byte
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 && j >= 0 {
		delta := int8(a[i]) - int8(b[j]) - borrow
		if delta < 0 {
			delta = 10 + delta
			borrow = 1
		}
		res = append(res, byte(delta)+'0')
		i--
		j--
	}

	for i >= 0 {
		delta := int8(a[i]) - '0' - borrow
		if delta < 0 {
			delta = 10 + delta
			borrow = 1
		} else {
			borrow = 0
		}
		res = append(res, byte(delta)+'0')
		i--
	}

	for j >= 0 {
		delta := int8(b[j]) - '0' - borrow
		if delta < 0 {
			delta = 10 + delta
			borrow = 1
		} else {
			borrow = 0
		}
		res = append(res, byte(delta)+'0')
		j--
	}

	for left, right := 0, len(res)-1; left < right; left, right = left+1, right-1 {
		res[left], res[right] = res[right], res[left]
	}
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return string(res)
}
