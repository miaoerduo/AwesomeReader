package middleware

import "github.com/PuerkitoBio/goquery"

type Middleware interface {
	Process(doc *goquery.Selection) *goquery.Selection
}
