package homework2

import "errors"

func PositionSwitchEncryptor(m string, key []int) (string, error) {
	if len(m)%len(key) != 0 {
		return "", errors.New("the length of the message must be a multiple of the length " +
			"of the key")
	}
	ms := []rune(m)
	cs := make([]rune, len(ms))
	p, t := 0, 0
	for i := 0; i < len(ms); i++ {
		cs[i] = ms[p*len(key)+key[t]-1]
		t++
		if t >= len(key) {
			t = 0
			p++
		}
	}
	return string(cs), nil
}
