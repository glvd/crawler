package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	db "github.com/bb/crawler/db"
	crawler "github.com/bb/crawler/lib/crawl"
	schema "github.com/bb/crawler/schema"
	"github.com/globalsign/mgo/bson"
)

func main() {
	c := new(crawler.Crawl)
	fmt.Println("****start crawl****")
	for page := 1; ; page++ {
		pageItems, err := c.CrawlPage(page)
		fmt.Println("<crawling page：", page, " items: ", len(pageItems), ">")
		if err != nil {
			fmt.Println("<--error-->", err.Error())
			failedLog(err.Error(), "page", strconv.Itoa(page))
			time.Sleep(1 * time.Minute)
			continue
		}

		for _, item := range pageItems {
			video := &schema.Video{}
			session := db.CloneSession()
			collection := session.DB("bus").C("videos")
			collection.Find(bson.M{"no": item.No}).One(&video)
			if video.No == item.No {
				fmt.Println("<crawled page: ", page, " no: ", item.No, ">")
				continue
			}
			fmt.Println("<----crawling no: ", item.No, " ---->")
			detail, err := c.CrawlDetail(item.No, item.Thumb)
			if err != nil {
				fmt.Println("<--error-->", err.Error())
				failedLog(err.Error(), "detail", item.No)
				time.Sleep(30 * time.Second)
				continue
			}

			collection.Insert(detail)
			session.Close()
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}

		if len(pageItems) < 30 {
			break
		}
	}
	os.Exit(1)
}

func failedLog(reason string, part string, no string) {
	session := db.CloneSession()
	defer session.Close()
	collection := session.DB("bus").C("failed")
	collection.Insert(&schema.Failed{Reason: reason, Part: part, No: no})
}
