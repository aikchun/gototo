package gototo_test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/aikchun/gototo"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestParseSelectionToDraw(t *testing.T) {
	data, err := os.ReadFile("draws_test.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))

	if err != nil {
		log.Fatal(err)
	}

	d := gototo.ParseSelectionToDraw(doc.Selection.Find("body").First())

	if d.GetDate() != "Thu, 09 Mar 2023" {
		t.Fatalf(`Did not get the correct date %v`, "Thu, 09 Mar 2023")
	}

	expectedWinningNumbers := []int{1, 10, 30, 31, 38, 45}

	if !Equal(d.GetWinningNumbers(), expectedWinningNumbers) {
		t.Fatalf(`Failed. Expected %v, instead got: %v`, expectedWinningNumbers, d.GetWinningNumbers())
	}

	expectedAdditionalNumber := 11

	if d.GetAdditionalNumber() != expectedAdditionalNumber {
		t.Fatalf(`Failed. Expected %v, instead got: %v`, expectedAdditionalNumber, d.GetAdditionalNumber())
	}

}

func TestParseNextDraw(t *testing.T) {
	data, err := os.ReadFile("nextdraw_test.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))

	if err != nil {
		log.Fatal(err)
	}

	d := gototo.ParseNextDraw(doc.Selection.Find("body").First())

	expectedDate := "Mon, 13 Mar 2023 , 6.30pm"

	if d.GetDate() != expectedDate {
		t.Fatalf(`Failed. Expected %s, instead got: %s`, expectedDate, d.GetDate())
	}

	expectedPrize := "$4,500,000 est"

	if d.GetPrize() != expectedPrize {
		t.Fatalf(`Failed. Expected %s, instead got: %s`, expectedPrize, d.GetPrize())
	}

}
