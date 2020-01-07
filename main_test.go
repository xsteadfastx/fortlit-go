package main

import (
	"testing"
)

func TestDecorate(t *testing.T) {
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
	}
	for _, table := range tables {
		r := table.q.decorate()
		if r != table.n {
			t.Errorf("string not is not \"%s\". got \"%s\".", table.n, r)
		}
	}
}