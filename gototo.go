package gototo

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func getDrawDate(t *goquery.Selection) string {
	return t.Find(".drawDate").Text()
}

func getWinningNumbers(t *goquery.Selection) []int {
	var s []int

	t.Find("td").Each(func(d int, td *goquery.Selection) {
		t := td.Text()
		i, _ := strconv.Atoi(t)
		s = append(s, i)
	})

	return s
}

func getAdditionalNumber(t *goquery.Selection) int {

	td := t.Find("td").First()
	text := td.Text()
	i, _ := strconv.Atoi(text)

	return i
}

func ParseSelectionToDraw(e *goquery.Selection) Draw {
	li := e.Find("li").Eq(0)
	tables := li.Find("table")

	date := getDrawDate(tables.Eq(0))
	winningNumbers := getWinningNumbers(tables.Eq(1))
	additionalNumber := getAdditionalNumber(tables.Eq(2))

	return DrawModel{
		Date:             date,
		WinningNumbers:   winningNumbers,
		AdditionalNumber: additionalNumber,
	}
}

func GetLatestDraw() Draw {

	c := colly.NewCollector()
	var d Draw

	c.OnHTML("body", func(e *colly.HTMLElement) {
		d = ParseSelectionToDraw(e.DOM)
	})

	c.Visit("https://www.singaporepools.com.sg/DataFileArchive/Lottery/Output/toto_result_top_draws_en.html")

	return d

}

func ParseNextDraw(e *goquery.Selection) NextDraw {

	return NextDrawModel{
		Date:  e.Find(".toto-draw-date").First().Text(),
		Prize: e.Find("span").First().Text(),
	}

}
func GetNextDraw() NextDraw {
	c := colly.NewCollector()

	var n NextDraw

	c.OnHTML("body", func(e *colly.HTMLElement) {
		n = ParseNextDraw(e.DOM)
	})

	c.Visit("https://www.singaporepools.com.sg/DataFileArchive/Lottery/Output/toto_next_draw_estimate_en.html")

	return n
}
