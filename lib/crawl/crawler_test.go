package crawler_test

import (
	"testing"

	crawler "github.com/crawler/lib/crawl"
)

const (
	no    = "JUFE-062"
	thumb = "https://pics.javbus.com/thumb/747o.jpg"
)

func TestCrawlDetail(t *testing.T) {
	c := new(crawler.Crawl)
	detail, err := c.CrawlDetail(no, thumb)
	if err != nil {
		t.Log("err", err)
	}
	t.Log(detail)
}
