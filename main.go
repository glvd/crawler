package main

import (
	crawler "github.com/bb/crawler/lib"
)

func main() {
	c := new(crawler.Crawl)
	c.CrawlDetail("DASD-539")
}
