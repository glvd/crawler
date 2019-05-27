package chromedp

import (
	"context"

	"github.com/chromedp/chromedp"
)

// LoadHTML ...
func LoadHTML(url string) (string, error) {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var html string

	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		// retrieve the value of the textarea
		chromedp.OuterHTML("body", &html),
	})
	if err != nil {
		return html, err
	}
	return html, err
}
