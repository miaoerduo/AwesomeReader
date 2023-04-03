package app

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/miaoerduo/AwesomeReader/app/middleware"
)

type Index struct {
	Id    string
	Label string
	Href  string
}

type EpubParser struct {
	// epub file path
	EpubPath  string
	zipReader *zip.ReadCloser
	// index of each page
	IndexList      []Index
	MiddleWareList []middleware.Middleware
}

func (app *EpubParser) Init() {
	var err error
	app.zipReader, err = zip.OpenReader(app.EpubPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range app.zipReader.File {
		if file.Name == "toc.ncx" {
			rc, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer rc.Close()
			app.GetIndex(rc)
		}
	}
}

func (app *EpubParser) Close() {
	app.zipReader.Close()
}

func (app *EpubParser) GetIndex(rc io.ReadCloser) {
	// parse toc.ncx
	doc, err := goquery.NewDocumentFromReader(rc)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("navMap navPoint").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("id")
		label, _ := s.Find("navLabel text").Html()
		href, _ := s.Find("content").Attr("src")
		app.IndexList = append(app.IndexList, Index{Id: id, Label: label, Href: href})
		fmt.Println(id, label, href)
	})
}

func (app *EpubParser) Dump(destDir string) {
	// index page
	indexPage := `<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<title>Index</title>
	</head>
	<body>
		<ul>
		{{range .}}
			<li><a href="{{.Href}}">{{.Label}}</a></li>
		{{end}}
		</ul>
	</body>
	</html>
	`
	// parse index page
	tmpl, err := template.New("index").Parse(indexPage)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(destDir + "/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, app.IndexList)
	if err != nil {
		log.Fatal(err)
	}
	// copy OEBPS files
	for _, file := range app.zipReader.File {
		if file.Name == "toc.ncx" {
			continue
		}
		if file.FileInfo().IsDir() {
			continue
		}
		if file.Name == "META-INF/container.xml" {
			continue
		}
		rc, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()
		dstFilePath := destDir + "/" + file.Name
		dstFolder := dstFilePath[0:strings.LastIndex(dstFilePath, "/")]
		if _, err := os.Stat(dstFolder); os.IsNotExist(err) {
			os.MkdirAll(dstFolder, 0755)
		}
		doc, _ := goquery.NewDocumentFromReader(rc)
		selection := doc.Selection
		for _, mw := range app.MiddleWareList {
			selection = mw.Process(selection)
		}

		f, err := os.Create(dstFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		html, _ := selection.Html()
		fmt.Fprint(f, html)
	}
}
