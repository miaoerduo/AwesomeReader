package app

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func TTSWrapp(htmlContent string) (string, error) {
	// 1. read a html page
	// 2. add a button on the top
	// 3. when click this button, it will call web browser to read the page via TTS
	// 4. when click any position, stop TTS, and record the position.
	// 5. when click the button again, it will read the page from the recorded position.

	doc := strings.NewReader(htmlContent)
	d, err := goquery.NewDocumentFromReader(doc)
	if err != nil {
		panic(err)
	}
	d.Find("body").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("Found %d %s", i, s.Text())
		s.AppendHtml(`<button onclick="window.speechSynthesis.speak(new SpeechSynthesisUtterance(document.body.innerText));">Read</button>`)
	})
	return d.Html()
}

// get all sound from web browser
// https://developer.mozilla.org/en-US/docs/Web/API/SpeechSynthesis
// https://developer.mozilla.org/en-US/docs/Web/API/SpeechSynthesisUtterance
// https://developer.mozilla.org/en-US/docs/Web/API/SpeechSynthesis/getVoices
// https://developer.mozilla.org/en-US/docs/Web/API/SpeechSynthesis/speak
