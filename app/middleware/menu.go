package middleware

import (
	"github.com/PuerkitoBio/goquery"
)

type Menu struct{}

func (menu *Menu) Process(doc *goquery.Selection) *goquery.Selection {
	body := doc.Find("body")
	body.AppendHtml(`
		<div id="menu-header" style="display: none; position: fixed; left: 0; top: 0; font-size: 40px; z-index: 100; color: gray; width: 100%; background: #ffffffe3; padding-left: 20; padding-right: 20;">
			<span id="index">☰</span>
			<span id="setting" style=" position: relative; float: right; right: 40;">⚙</span>
		</div>`)
	body.AppendHtml(`<div id="menu-tail" style="display: none; position: fixed; left: 0; bottom: 0; font-size: 40px; z-index: 100; color: gray; width: 100%; background: #ffffffe3; padding-left: 20; padding-right: 20;"><span id="menu-play">▷</span></div>`)

	// move script to the bottom
	script := doc.Find("script")
	script.Remove()
	body.AppendSelection(script)

	return doc
}
