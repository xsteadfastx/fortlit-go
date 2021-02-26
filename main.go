// nolint: exhaustivestruct, gochecknoglobals, gomnd
package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/pterm/pterm"
	"go.xsfx.dev/fortlit/quotes"
)

//go:generate go run tools/gen/gen.go
//go:generate gofumpt -w quotes/quotes.go

const lineWords = 8

var (
	ColorAuthor = pterm.NewRGB(189, 147, 249)
	ColorTime   = pterm.NewRGB(80, 250, 123)
	ColorBook   = pterm.NewRGB(255, 121, 198)
)

func getQuote(qs map[string][]quotes.Quote, t string) quotes.Quote {
	var q quotes.Quote

	if _, ok := qs[t]; ok {
		if len(qs[t]) != 1 {
			rand.Seed(time.Now().Unix())
			q = qs[t][rand.Intn(len(qs[t]))] // nolint: gosec
		} else {
			q = qs[t][0]
		}
	}

	return q
}

func stringWrap(text string, limit int) string {
	ts := strings.Fields(text)
	rs := ""

	wc := 0
	for _, i := range ts {
		if wc < limit {
			rs = rs + " " + i
			wc++
		} else {
			rs = rs + " " + i + "\n"
			wc = 0
		}
	}

	return rs
}

func main() {
	now := time.Now()
	t := fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())
	q := getQuote(quotes.FortData, t)

	if q == (quotes.Quote{}) {
		return
	}

	m := regexp.MustCompile(fmt.Sprintf("(?i)(%s)", q.Time))
	text := m.ReplaceAllString(q.Text, ColorTime.Sprint("$1"))

	pterm.DefaultCenter.Println(
		pterm.DefaultBox.Sprint(
			stringWrap(text, lineWords),
		),
	)
	pterm.DefaultCenter.Println(fmt.Sprintf("✍️ %s - 📖 %s", ColorAuthor.Sprint(q.Author), ColorBook.Sprint(q.Book)))
}
