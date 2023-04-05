package middleware

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Dict struct{}

func readContent(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func (dict *Dict) Process(doc *goquery.Selection) *goquery.Selection {
	doc.Find("head").AppendHtml(`<meta name="viewport" content="width=device-width, initial-scale=1">`)
	doc.Find("head").AppendHtml(`<meta charset="utf-8">`)

	_, filename, _, _ := runtime.Caller(0)

	dictJsPath := filename[:strings.LastIndex(filename, "/")] + "/dict.js"

	script := readContent(dictJsPath)
	body := doc.Find("body")
	body.WrapInnerHtml(`<div id="main-book"></div>`)
	body.AppendHtml("<script src=\"https://unpkg.com/sweetalert/dist/sweetalert.min.js\"></script>")
	body.AppendHtml("<script src=\"https://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js\"></script>")
	doc.Find("body").AppendHtml(`<script>` + script + `</script>`)

	return doc
}
