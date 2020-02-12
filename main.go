package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/xsteadfastx/fortlit-go/data"
)

//go:generate go-bindata -pkg data -o ./data/bindata.go litdata.json
//go:generate go fmt ./data/bindata.go

var version = "development"

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
		//nolint: gomnd
		if len(qs[t]) != 1 {
			rand.Seed(time.Now().Unix())
			quote = qs[t][rand.Intn(len(qs[t]))]
		} else {
			quote = qs[t][0]
		}
	}

	return quote
}

func open(as string) []byte {
	data, err := data.Asset(as)

	if err != nil {
		panic(err)
	}

	return data
}

func (q *Quote) decorate() string {
	m := regexp.MustCompile(fmt.Sprintf("(?i)(%s)", q.Time))
	text := m.ReplaceAllString(q.Text, fmt.Sprintf(Purple, "$1"))

	return fmt.Sprintf("\n%s\n\n    - %s, %s\n", text, q.Book, fmt.Sprintf(Teal, q.Author))
}

func main() {
	fversion := flag.Bool("version", false, "Shows version.")

	flag.Parse()

	if *fversion {
		fmt.Printf("Version: %s", version)
		os.Exit(0)
	}

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
