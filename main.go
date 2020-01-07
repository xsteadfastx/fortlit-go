package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:generate go-bindata litdata.json
//go:generate go fmt bindata.go

const (
	Purple = "\033[1;34m%s\033[0m"
	Teal   = "\033[1;36m%s\033[0m"
)

type Quote struct {
	Author string `json:"author"`
	Book   string `json:"book"`
	Text   string `json:"text"`
	Time   string `json:"time"`
}

func get(qs map[string][]Quote, t string) Quote {
	quote := Quote{}

	if _, exists := qs[t]; exists {
		if len(qs[t]) != 1 {
			rand.Seed(time.Now().Unix())
			quote = qs[t][rand.Intn(len(qs[t]))]
		}
	}

	return quote

}

func open(as string) []byte {
	data, err := Asset(as)

	if err != nil {
		panic(err)
	}

	return data
}

func (q *Quote) decorate() string {
	text := strings.Replace(q.Text, q.Time, fmt.Sprintf(Purple, q.Time), -1)

	return fmt.Sprintf("\n%s\n\n    - %s, %s\n", text, q.Book, fmt.Sprintf(Teal, q.Author))
}

func main() {
	data := open("litdata.json")
	qs := make(map[string][]Quote)
	if err := json.Unmarshal(data, &qs); err != nil {
		panic(err)
	}
	now := time.Now()
	t := fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())
	quote := get(qs, t)
	if quote != (Quote{}) {
		fmt.Println(quote.decorate())
	}
}
