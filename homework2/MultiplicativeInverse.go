package homework2

import (
	"errors"
)

func exgcd(a int, b int) (int, int, int) {
	if b == 0 {
		return 1, 0, a
	} else {
		x, y, q := exgcd(b, a % b)
		return y, x - y * (a/b), q
	}
}

func gcd(a int, b int) int {
	if a < b {
		a, b = b, a
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return b
}

func MultiplicativeInverse(num int, mod int) (int, error) {
	x, _, q := exgcd(num, mod)
	if q != 1 {
		return 0, errors.New("gcd(MOD, num) != 1")
	}
	return ((x % mod) + mod) % mod, nil
}
