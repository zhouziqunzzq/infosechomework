package homework1

import (
	"sort"
	"math"
)

type column struct {
	index rune
	cid   int
}

func KeyControlledEncryptor(message string, key string) (string) {
	keys := []rune(key)
	messages := []rune(message)
	columnCount := len(key)
	rowCount := int(math.Ceil(float64(len(message)) / float64(columnCount)))
	columns := make([]column, 0)
	for i := 0; i < columnCount; i++ {
		columns = append(columns, column{index: keys[i], cid: i})
	}
	sort.SliceStable(columns, func(i, j int) bool {
		return columns[i].index < columns[j].index
	})
	crypted := make([]rune, 0)
	t := 0
	for i := 0; i < columnCount; i++ {
		for j := 0; j < rowCount; j++ {
			t = columnCount*j + columns[i].cid
			if t >= len(message) {
				crypted = append(crypted, ' ')
			} else {
				crypted = append(crypted, messages[t])
			}
		}
	}
	return string(crypted)
}
