package app

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func DictWrapper(htmlContent string) (string, error) {
	// when TOUCH any word (NOT SELECT)
	//  call dict to get the meaning of the word, the api is https://api.dictionaryapi.dev/api/v2/entries/en/{word}
	// and show the meaning on the page.

	doc := strings.NewReader(htmlContent)
	d, err := goquery.NewDocumentFromReader(doc)
	if err != nil {
		panic(err)
	}
	d.Find("body").Each(func(i int, s *goquery.Selection) {
		// each word should in a span
		s.Find("p").Each(func(i int, s *goquery.Selection) {
			if s.Text() == "" {
				return
			}
			c, _ := s.Children().Html()
			s.SetHtml(strings.ReplaceAll(c, " ", `</span> <span>`))
		})

		s.AppendHtml(`<script>
		document.addEventListener('touchstart', function(event) {
			// get keyword (NOT all Text) from touch point
			var word = document.elementFromPoint(event.touches[0].clientX, event.touches[0].clientY).innerText;
			if (word) {
				var url = 'https://api.dictionaryapi.dev/api/v2/entries/en/' + word;
				var xhr = new XMLHttpRequest();
				xhr.open('GET', url, true);
				xhr.onload = function() {
					if (xhr.status === 200) {
						var result = JSON.parse(xhr.responseText);
						if (result.length > 0) {
							var meaning = result[0].meanings[0].definitions[0].definition;
							alert(word + "|" + meaning);
						}
					}
				};
				xhr.send();
			}
		});
		</script>`)
	})
	return d.Html()
}
