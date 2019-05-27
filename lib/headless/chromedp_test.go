package chromedp_test

import (
	"testing"

	chromedp "github.com/bus_crawler/lib/headless"
)

const (
	url = "https://www.seedmm.life"
)

func TestHeadless(t *testing.T) {
	html, err := chromedp.LoadHTML(url)
	if err != nil {
		t.Log(err)
	}
	t.Log(html)
}
