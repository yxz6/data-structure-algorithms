package leetcode

func power(x float64, n int) float64 {
	/*
		Zero to the power of zero, denoted by 00,
		is a mathematical expression that is either defined as 1 or left undefined, depending on context.
		In algebra and combinatorics, one typically defines  00 = 1. In mathematical analysis, the expression is sometimes left undefined. Computer programming languages and software also have differing ways of handling this expression.
	*/

	if n == 0 {
		return 1
	}

	if x == 0 {
		return 0
	}

	if n < 0 {
		return 1 / power(x, -n)
	}

	if n == 1 {
		return x
	}

	k := power(x, n/2)
	k *= k
	if n%2 == 1 {
		k *= x
	}
	return k
}
