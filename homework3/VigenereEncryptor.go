package homework3

import (
	"strings"
)

func VigenereEncryptor(m string, key string) string {
	// Convert message to upper cases first
	ms := []rune(strings.ToUpper(m))
	cs := make([]rune, len(ms))
	ks := []rune(strings.ToUpper(key))
	for i := 0; i < len(ms); i++ {
		cs[i] = rune('A' + (ms[i]-'A'+ks[i%len(ks)]-'A')%26)
	}
	return string(cs)
}
