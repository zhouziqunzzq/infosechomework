package homework1

func DoubleRailEncryptor(m string) (string) {
	ms := []rune(m)
	cs := make([]rune, len(m))
	p := 0
	for i := 0; i < len(m); i += 2 {
		cs[p] = ms[i]
		p++
	}
	for i := 1; i < len(m); i += 2 {
		cs[p] = ms[i]
		p++
	}
	return string(cs)
}
