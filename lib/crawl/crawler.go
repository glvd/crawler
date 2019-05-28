package crawler

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/anaskhan96/soup"
	schema "github.com/crawler/schema"
)

const (
	url = "https://www.seedmm.life"
)

// Crawl ...
type Crawl struct {
}

// PageItems ...
type PageItems struct {
	URL  string
	Name string
}

// ListItems ...
type ListItems struct {
	No    string
	Thumb string
	Title string
}

// CrawlActress ...
func (c *Crawl) CrawlActress(aURL string, page int) ([]ListItems, error) {
	listURL := fmt.Sprintf("%s/%d", aURL, page)
	resp, err := soup.Get(listURL)
	if err != nil {
		return nil, err
	}
	doc := soup.HTMLParse(resp)
	infos := doc.Find("div", "id", "waterfall").FindAll("a", "class", "movie-box")
	list := []ListItems{}
	for _, info := range infos {
		thumbInfo := info.Find("div", "class", "photo-frame").Find("img").Attrs()
		item := ListItems{
			Thumb: thumbInfo["src"],
			No:    info.Find("div", "class", "photo-info").Find("date").Text(),
			Title: info.Find("div", "class", "photo-info").Find("span").Text(),
		}
		list = append(list, item)
	}
	return list, nil
}

// CrawlPage ...
func (c *Crawl) CrawlPage(page int, mode string) ([]PageItems, error) {
	items := []PageItems{}
	var actressURL string
	if mode == "2" {
		actressURL = fmt.Sprintf("%s/uncensored/actresses/%d", url, page)
	} else {
		actressURL = fmt.Sprintf("%s/actresses/%d", url, page)
	}
	resp, err := soup.Get(actressURL)
	if err != nil {
		return items, err
	}
	doc := soup.HTMLParse(resp)
	infos := doc.Find("div", "id", "waterfall").FindAll("div", "class", "item")
	for _, info := range infos {
		link := info.Find("a").Attrs()
		name := info.Find("div", "class", "photo-info").Find("span").Text()
		pageItem := PageItems{link["href"], name}

		items = append(items, pageItem)
	}

	return items, nil
}

// CrawlDetail ...
func (c *Crawl) CrawlDetail(no string, thumb string, title string) (*schema.Video, error) {
	var err error
	video := &schema.Video{}
	detailURL := fmt.Sprintf("%s/%s", url, no)
	soup.Get(detailURL)
	res, err := soup.Get(detailURL)
	if err != nil {
		return video, err
	}

	doc := soup.HTMLParse(res)
	details := doc.Find("div", "class", "container").Find("div", "class", "movie")
	coverInfo := details.Find("div", "class", "screencap").Find("a").Attrs()
	infos := details.Find("div", "class", "info").Children()

	video.Thumb, err = getImg(thumb)
	video.Cover, err = getImg(coverInfo["href"])
	video.Title = title
	video.Tags = buildDesc(details.FindAll("span", "class", "genre"))
	video.Stars = buildDesc(details.FindAll("div", "class", "star-name"))
	for _, info := range infos {
		labelMatch(info, video)
	}
	return video, err
}

func getImg(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	encodeString := base64.StdEncoding.EncodeToString(pix)
	return fmt.Sprintf("data:image/png;base64,%s", encodeString), err
}

func buildDesc(v []soup.Root) []string {
	var descs []string
	for _, tag := range v {
		text := tag.Find("a").Text()
		descs = append(descs, text)
	}
	return descs
}

func labelMatch(dom soup.Root, video *schema.Video) {
	texts := dom.FullText()
	t := strings.Split(texts, ":")
	if len(t) == 2 {
		field := strings.Replace(strings.TrimSpace(t[0]), "\n", "", -1)
		value := strings.Replace(strings.TrimSpace(t[1]), "\n", "", -1)
		switch field {
		case "識別碼":
			video.No = value
		case "發行日期":
			video.Date = value
		case "長度":
			video.Length = value
		case "導演":
			video.Director = value
		case "製作商":
			video.Producer = value
		case "發行商":
			video.Publisher = value
		case "系列":
			video.Series = value
		}
	}

}
