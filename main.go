package main

import (
	"strconv"
	"time"

	db "github.com/bb/crawler/db"
	crawler "github.com/bb/crawler/lib/crawl"
	schema "github.com/bb/crawler/schema"
	"github.com/globalsign/mgo/bson"
)

func main() {
	c := new(crawler.Crawl)
	for page := 1; ; page++ {
		pageItems, err := c.CrawlPage(page)
		if err != nil {
			failedLog(err.Error(), "page", strconv.Itoa(page))
			time.Sleep(1 * time.Minute)
			continue
		}

		for _, item := range pageItems {
			video := &schema.Video{}

			session := db.CloneSession()
			defer session.Close()
			collection := session.DB("bus").C("videos")
			collection.Find(bson.M{"no": item.No}).One(&video)

			if video.No == item.No {
				continue
			}

			detail, err := c.CrawlDetail(item.No, item.Thumb)
			if err != nil {
				failedLog(err.Error(), "detail", item.No)
				time.Sleep(30 * time.Second)
				continue
			}

			collection.Insert(detail)
			time.Sleep(5 * time.Second)
		}

		if len(pageItems) < 30 {
			break
		}
		time.Sleep(10 * time.Second)
	}
}

func failedLog(reason string, part string, no string) {
	session := db.CloneSession()
	defer session.Close()
	collection := session.DB("bus").C("failed")
	collection.Insert(&schema.Failed{Reason: reason, Part: part, No: no})
}
