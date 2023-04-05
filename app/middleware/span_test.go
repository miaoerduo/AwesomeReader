package middleware

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestSpan(t *testing.T) {
	doc := `
	<body>
	<p class="chaptersubtitle"><span class="bold">Y<span class="smallcaps">ALI’S</span> Q<span class="smallcaps">UESTION</span></span></p
	</body>
	`
	html, _ := goquery.NewDocumentFromReader(strings.NewReader(doc))

	span := &Span{}

	span.Process(html.Selection)

	s, _ := html.Html()
	fmt.Printf("html == %+v\n\n", s)
}

func TestRune(t *testing.T) {
	s := "Still other hunter-gatherers in contact with farmers did eventually become farmers, but only after what may seem to us like an inordinately long delay. For example, the coastal peoples of northern Germany did not adopt food production until 1,300 years after peoples of the Linearbandkeramik culture introduced it to inland parts of Germany only 125 miles to the south. Why did those coastal Germans wait so long, and what led them finally to change their minds?"
	word := ""
	for _, symbol := range s {
		if IsSymbol(symbol) {
			if word != "" {
				fmt.Println(word)
				word = ""
			}
			fmt.Println(string(symbol))
			continue
		}
		word += string(symbol)
	}
	fmt.Println(word)
}
