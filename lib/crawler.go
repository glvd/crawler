package crawler

import (
	"fmt"

	"github.com/anaskhan96/soup"
)

const (
	url = "https://www.seedmm.life"
)

// Crawl ...
type Crawl struct {
}

// ListItems ...
type ListItems struct {
	no    string
	thumb string
}

// CrawlPage ...
func (c *Crawl) CrawlPage(page int) []ListItems {
	listURL := fmt.Sprintf("%s/page/%d", url, page)
	resp, err := soup.Get(listURL)
	if err != nil {
		fmt.Print(err)
	}
	doc := soup.HTMLParse(resp)
	infos := doc.Find("div", "id", "waterfall").FindAll("a")
	list := []ListItems{}
	for _, info := range infos {
		videoInfo := info.Find("div", "class", "photo-frame").Find("img").Attrs()
		item := ListItems{
			thumb: videoInfo["src"],
			no:    info.Find("div", "class", "photo-info").Find("date").Text(),
		}
		list = append(list, item)
	}
	return list
}

// CrawlDetail ...
func (c *Crawl) CrawlDetail(no string) {
	// v := &schema.Video{}
	detailURL := fmt.Sprintf("%s/%s", url, no)
	soup.Get(detailURL)
	res, err := soup.Get(detailURL)
	if err != nil {
		fmt.Print(err)
	}
	doc := soup.HTMLParse(res)

	details := doc.Find("div", "class", "container").FindAll("div", "class", "movie")
	// coverInfo := detail.Find("div", "class", "screencap").Find("a").Attrs()
	infos := details[0].Find("div", "class", "info").Children()
	stars := details[0].FindAll("div", "class", "star-name")
	magnetLinks := doc.Find("table", "id", "magnet-table").FindAll("tr")
	fmt.Print(magnetLinks)
	// for _, link := range magnetLinks {
	// 	fmt.Print(link)
	// }
	for _, star := range stars {
		fmt.Print(star.Find("a").Text())
	}
	for _, info := range infos {
		children := info.Children()
		for _, child := range children {
			label := child.Text()

			fmt.Print(labelMatch(label))
		}
		fmt.Print(info.Text())
	}
}

func labelMatch(label string) string {
	switch label {
	case "識別碼:":
		return "no"
	case "發行日期:":
		return "date"
	case "長度:":
		return "length"
	case "導演:":
		return "director"
	case "製作商:":
		return "producer"
	case "發行商:":
		return "publisher"
	case "系列:":
		return "series"
	case "類別:":
		return "tag"
	case "演員:":
		return "actress"
	}
	return label
}
