package middleware

import (
	"github.com/PuerkitoBio/goquery"
)

type Dict struct{}

func (dict *Dict) Process(doc *goquery.Selection) *goquery.Selection {
	doc.Find("head").AppendHtml(`<meta name="viewport" content="width=device-width, initial-scale=1">`)
	doc.Find("head").AppendHtml(`<meta charset="utf-8">`)

	script := `
	function checkDict(elem) {
		if (elem.className !== 'med-word') {
			return
		}
		word = elem.innerHTML;
		word = word.replace(/[^\w\s]/gi, '');
		if (!word) {
			return
		}
		var url = 'https://api.dictionaryapi.dev/api/v2/entries/en/' + word;
		var xhr = new XMLHttpRequest();
		xhr.open('GET', url, true);
		xhr.onload = function () {
			if (xhr.status === 200) {
				var result = JSON.parse(xhr.responseText);
				if (result.length > 0) {
					var meaning = result[0].meanings[0].definitions[0].definition;
					swal(word, meaning)
				}
			}
		};
		xhr.send();
	}
	
	document.addEventListener('click', function (event) {
		if (event.target.className != 'med-word') {
			return
		}
		let elem = document.elementFromPoint(event.clientX, event.clientY);
		checkDict(elem)
	});	
	`
	body := doc.Find("body")
	body.AppendHtml("<script src=\"https://unpkg.com/sweetalert/dist/sweetalert.min.js\"></script>")
	doc.Find("body").AppendHtml(`<script>` + script + `</script>`)

	return doc
}
