package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	db "github.com/crawler/db"
	crawler "github.com/crawler/lib/crawl"
	schema "github.com/crawler/schema"
	"github.com/globalsign/mgo/bson"
)

var (
	method       string
	mode         string
	inputBangumi string
	skipCount    = 0
	count        = 0
)

func main() {
	// choose run mode
	fmt.Println("选择运行方式：1.全量 2.查询 默认：全量")
	fmt.Scanln(&method)
	if method == "2" {
		fmt.Println("请输入要爬的番号，以逗号隔开: ")
		fmt.Scanln(&inputBangumi)
	} else {
		fmt.Println("选择爬虫模式: 1.有码 2.无码 默认：有码")
		fmt.Scanln(&mode)
	}

	// err process
	defer func() {
		if err := recover(); err != nil {
			log.Panic("<--panic error-->", err)
		}
	}()

	// init Crawler
	c := new(crawler.Crawl)

	fmt.Println("****start crawl****")
	if method == "2" {
		inputs := strings.Split(inputBangumi, ",")
		for _, input := range inputs {
			log.Println("searching bangumi", input)
			items, err := c.Search(input)
			if err != nil {
				fmt.Println("<--error-->", err.Error())
				failedLog(err.Error(), "search", input)
				continue
			}

			for _, item := range items {
				checkRes := checkCrawled(item.No)

				fmt.Println("<crawling bangumi: ", item.No, ">")
				detail, err := c.CrawlDetail(item.No, item.Thumb, item.Title)
				if err != nil {
					fmt.Println("<--error-->", err.Error())
					failedLog(err.Error(), "detail", item.No)
					if err.Error() == "invalid memory address or nil pointer dereference" {
						break
					} else {
						continue
					}

				}
				detail.Uncensored = false
				if checkRes == false {
					createRecord(detail)
				}

				time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			}
		}
	} else {
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
	}
	fmt.Println("完成啦")
}

func crawlActress(c *crawler.Crawl, actress crawler.PageItems) {
	url := actress.URL
	name := actress.Name
	fmt.Println("<crawling actress: ", name, ">")

	for page := 1; ; page++ {
		if skipCount >= 20 {
			skipCount = 0
			break
		}
		pageItems, err := c.CrawlActress(url, page)
		if err != nil {
			fmt.Println("<--error-->", err.Error())
			failedLog(err.Error(), "page", strconv.Itoa(page))
			if err.Error() == "invalid memory address or nil pointer dereference" {
				break
			} else {
				continue
			}

		}
		for _, item := range pageItems {
			checkRes := checkCrawled(item.No)
			if skipCount >= 20 {
				fmt.Println("<crawling actress: ", name, " skip>")
				break
			}
			if checkRes == true {
				fmt.Println("<crawled actress: ", name, "no: ", item.No, ">")
				skipCount++
				continue
			}
			fmt.Println("<crawling actress: ", name, "no: ", item.No, ">")
			detail, err := c.CrawlDetail(item.No, item.Thumb, item.Title)
			if err != nil {
				fmt.Println("<--error-->", err.Error())
				failedLog(err.Error(), "detail", item.No)
				if err.Error() == "invalid memory address or nil pointer dereference" {
					break
				} else {
					continue
				}

			}
			if mode == "2" {
				detail.Uncensored = true
			} else {
				detail.Uncensored = false
			}
			createRecord(detail)
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		}
		if len(pageItems) < 30 {
			break
		}
	}
}

func createRecord(video *schema.Video) bool {
	session := db.CloneSession()
	defer session.Close()
	collection := session.DB("bus").C("videos")
	collection.Insert(video)
	return true
}

func checkCrawled(no string) bool {
	session := db.CloneSession()
	defer session.Close()
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
