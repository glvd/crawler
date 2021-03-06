package crawler_test

import (
	"testing"

	crawler "github.com/crawler/crawler/lib/crawl"
)

const (
	no    = "JUFE-062"
	thumb = "https://pics.javbus.com/thumb/747o.jpg"
)

func TestCrawlPage(t *testing.T) {
	c := new(crawler.Crawl)
	res, err := c.CrawlPage(1, "1")
	t.Log(len(res), err)
}

func TestCrawlActress(t *testing.T) {
	c := new(crawler.Crawl)

	urls, err := c.CrawlActress("https://www.javbus.com/star/okq", 1)
	t.Log(len(urls), err)
}
