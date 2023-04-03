package middleware

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type Span struct {
	textNodeList []*html.Node
}

func (span *Span) Process(doc *goquery.Selection) *goquery.Selection {
	// find all text node
	for _, node := range doc.Find("body").Nodes {
		span.FindTextNode(node)
	}
	for _, node := range span.textNodeList {
		span.AddSpan(node)
	}
	return doc
}

func (span *Span) FindTextNode(node *html.Node) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		span.FindTextNode(c)
	}
	if node.Type == html.TextNode {
		span.textNodeList = append(span.textNodeList, node)
	}
}

func IsSymbol(r rune) bool {
	return r == ' ' || r == ',' || r == '.' || r == '!' || r == '?' || r == ';' || r == ':' || r == '"' || r == '\'' || r == '(' || r == ')' || r == '[' || r == ']' || r == '{' || r == '}' || r == '\n' || r == '\t'
}

// add span tag for each word
func (span *Span) AddSpan(node *html.Node) {
	word := ""
	parent := node.Parent
	if parent == nil {
		return
	}
	parent.RemoveChild(node)
	for _, symbol := range node.Data {
		if IsSymbol(symbol) {
			if word != "" {
				spanNode := &html.Node{
					Type: html.ElementNode,
					Data: "span",
					Attr: []html.Attribute{
						{
							Key: "class",
							Val: "med-word",
						},
					},
				}
				spanNode.AppendChild(&html.Node{
					Type: html.TextNode,
					Data: word,
				})
				parent.AppendChild(spanNode)
				word = ""
			}
			parent.AppendChild(&html.Node{
				Type: html.TextNode,
				Data: string(symbol),
			})
			continue
		}
		word += string(symbol)
	}
	if word != "" {
		spanNode := &html.Node{
			Type: html.ElementNode,
			Data: "span",
			Attr: []html.Attribute{
				{
					Key: "class",
					Val: "med-word",
				},
			},
		}
		spanNode.AppendChild(&html.Node{
			Type: html.TextNode,
			Data: word,
		})
		parent.AppendChild(spanNode)
	}
}
