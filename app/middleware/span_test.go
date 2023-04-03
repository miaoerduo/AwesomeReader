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
	<p class="chapteropenertext"><span class="chapteropenerfirstletters">F</span><span class="smallcaps1">ORMERLY, ALL PEOPLE ON EARTH WERE HUNTER-GATHERERS</span>. Why did any of them adopt food production at all? Given that they must have had some reason, why did they do so around 8500 <span class="smallcaps1">B.C.</span> in Mediterranean habitats of the Fertile Crescent, only 3,000 years later in the climatically and structurally similar Mediterranean habitats of southwestern Europe, and never indigenously in the similar Mediterranean habitats of California, southwestern Australia, and the Cape of South Africa? Why did people of the Fertile Crescent wait until 8500 <span class="smallcaps1">B.C.</span>, instead of becoming food producers around 18,500 or 28,500 <span class="smallcaps1">B.C.</span>?</p>

	<p class="para">From our modern perspective, all these questions at first seem silly, because the drawbacks of being a hunter-gatherer appear so obvious. Scientists used to quote a phrase of Thomas Hobbes’s in order to characterize the lifestyle of hunter-gatherers as “nasty, brutish, and short.” They seemed to have to work hard, to be driven by the daily quest for food, often to be close to starvation, to lack such elementary material comforts as soft beds and adequate clothing, and to die young.</p>

	<p class="para">In reality, only for today’s affluent First World citizens, who don’t actually do the work of raising food themselves, does food production (by remote agribusinesses) mean less physical work, more comfort, freedom from starvation, and a longer expected lifetime. Most peasant farmers and herders, <a id="page_101" class="calibre4"></a>who constitute the great majority of the world’s actual food producers, aren’t necessarily better off than hunter-gatherers. Time budget studies show that they may spend more rather than fewer hours per day at work than hunter-gatherers do. Archaeologists have demonstrated that the first farmers in many areas were smaller and less well nourished, suffered from more serious diseases, and died on the average at a younger age than the hunter-gatherers they replaced. If those first farmers could have foreseen the consequences of adopting food production, they might not have opted to do so. Why, unable to foresee the result, did they nevertheless make that choice?</p>

	<p class="para">There exist many actual cases of hunter-gatherers who did see food production practiced by their neighbors, and who nevertheless refused to accept its supposed blessings and instead remained hunter-gatherers. For instance, Aboriginal hunter-gatherers of northeastern Australia traded for thousands of years with farmers of the Torres Strait Islands, between Australia and New Guinea. California Native American hunter-gatherers traded with Native American farmers in the Colorado River valley. In addition, Khoi herders west of the Fish River of South Africa traded with Bantu farmers east of the Fish River, and continued to dispense with farming themselves. Why?</p>

	<p class="para">Still other hunter-gatherers in contact with farmers did eventually become farmers, but only after what may seem to us like an inordinately long delay. For example, the coastal peoples of northern Germany did not adopt food production until 1,300 years after peoples of the Linearbandkeramik culture introduced it to inland parts of Germany only 125 miles to the south. Why did those coastal Germans wait so long, and what led them finally to change their minds?</p>

	<p class="spacebreak"> </p>
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
