package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	db "github.com/crawler/db"
	crawler "github.com/crawler/lib/crawl"
	schema "github.com/crawler/schema"
	"github.com/globalsign/mgo/bson"
)

var (
	mode      string
	skipCount = 0
)

func main() {
	fmt.Println("选择爬虫模式: 1.有码 2.无码 默认：有码")
	fmt.Scanln(&mode)
	c := new(crawler.Crawl)
	fmt.Println("****start crawl****")
	for page := 1; ; page++ {
		actresses, err := c.CrawlPage(page, mode)
		fmt.Println("<crawling page：", page, " items: ", len(actresses), ">")
		if err != nil {
			fmt.Println("<--error-->", err.Error())
			failedLog(err.Error(), "page", strconv.Itoa(page))
			time.Sleep(15 * time.Second)
			continue
		}
		for _, actress := range actresses {
			crawlActress(c, actress)
		}
		if len(actresses) < 50 {
			break
		}
	}
	os.Exit(1)
}

func crawlActress(c *crawler.Crawl, actress crawler.PageItems) {
	url := actress.URL
	name := actress.Name
	session := db.CloneSession()
	defer session.Clone()
	fmt.Println("<crawling actress: ", name, ">")

	for page := 1; ; page++ {
		pageItems, err := c.CrawlActress(url, page)
		if err != nil {
			fmt.Println("<--error-->", err.Error())
			failedLog(err.Error(), "page", strconv.Itoa(page))
			time.Sleep(10 * time.Second)
			continue
		}
		for _, item := range pageItems {
			checkRes := checkCrawled(item.No)
			if checkRes == true {
				fmt.Println("<crawled actress: ", name, "no: ", item.No, ">")
				skipCount++
				continue
			}
			fmt.Println("<----crawling no: ", item.No, " ---->")
			detail, err := c.CrawlDetail(item.No, item.Thumb, item.Title)
			if err != nil {
				fmt.Println("<--error-->", err.Error())
				failedLog(err.Error(), "detail", item.No)
				time.Sleep(30 * time.Second)
				continue
			}
			collection := session.DB("bus").C("videos")
			collection.Insert(detail)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
		if len(pageItems) < 30 {
			break
		}
	}
}

func checkCrawled(no string) bool {
	session := db.CloneSession()
	defer session.Clone()
	video := &schema.Video{}
	collection := session.DB("bus").C("videos")
	collection.Find(bson.M{"no": no}).One(&video)
	if video.No == no {
		return true
	}
	return false
}

func failedLog(reason string, part string, no string) {
	session := db.CloneSession()
	defer session.Close()
	collection := session.DB("bus").C("failed")
	collection.Insert(&schema.Failed{Reason: reason, Part: part, No: no})
}
