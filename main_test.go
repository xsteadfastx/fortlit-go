package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestGenWordSlice(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		input    []string
		index    int
		expected [][]string
	}{
		{
			[]string{"foo", "bar", "zonk"},
			2,
			[][]string{{"foo", "bar"}, {"zonk"}},
		},
	}

	for _, table := range tables {
		assert.Equal(genWordSlice(table.input, table.index), table.expected)
	}
}

// func TestGet(t *testing.T) {
// 	tables := []struct {
// 		m map[string][]Quote
// 		q Quote
// 	}{
// 		{
// 			map[string][]Quote{"00:00": {
// 				Quote{"Max Mustermann", "Testbook", "This is a time!", "time"},
// 			}},
// 			Quote{"Max Mustermann", "Testbook", "This is a time!", "time"},
// 		},
// 	}
// 	for _, table := range tables {
// 		q := get(table.m, "00:00")
// 		if q != table.q {
// 			t.Errorf("quote \"%+v\" is not like \"%+v\"", q, table.q)
// 		}
// 	}
// }
