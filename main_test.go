package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorate(t *testing.T) {
	assert := assert.New(t)
	tables := []struct {
		q Quote
		n string
	}{
		{
			Quote{
				"Max Mustermann",
				"Testbook",
				"This is a time!",
				"time",
			},
			"\nThis is a \033[1;34mtime\033[0m!\n\n    - Testbook, \033[1;36mMax Mustermann\033[0m\n",
		},
		{
			Quote{
				"Iain Banks",
				"Espedair Street",
				"Three twenty-three! Is that all?",
				"three twenty-three",
			},
			"\n\033[1;34mThree twenty-three\033[0m! Is that all?\n\n    - Espedair Street, \033[1;36mIain Banks\033[0m\n",
		},
	}

	for _, table := range tables {
		r := table.q.decorate()
		assert.Equal(table.n, r)
	}
}

func TestGet(t *testing.T) {
	tables := []struct {
		m map[string][]Quote
		q Quote
	}{
		{
			map[string][]Quote{"00:00": {
				Quote{"Max Mustermann", "Testbook", "This is a time!", "time"},
			}},
			Quote{"Max Mustermann", "Testbook", "This is a time!", "time"},
		},
	}
	for _, table := range tables {
		q := get(table.m, "00:00")
		if q != table.q {
			t.Errorf("quote \"%+v\" is not like \"%+v\"", q, table.q)
		}
	}
}
