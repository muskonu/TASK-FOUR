package main

import (
	"fmt"
	"ggggggggggggo/pkg"
	"github.com/gocolly/colly"
	"strconv"
)

func Crawl(url string) {
	c := colly.NewCollector()
	var books []pkg.Book
	c.OnHTML("div.channel-item", func(e *colly.HTMLElement) {
		sbook := pkg.SeBook{}
		err := e.Unmarshal(&sbook)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		book := pkg.Book{}
		book.SeBook = sbook
		books = append(books, book)
	})
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Visit error:", err)
		return
	}
	db := pkg.CreateDb()
	db.Create(&books)
} //爬取该页面的作者，书名以及分类

func CrawlPage(i int, page chan<- int) {
	var url string
	if i == 1 {
		url = "https://kgbook.com/shigesanwen/"
	} else {
		url = "https://kgbook.com/shigesanwen/index_" + strconv.Itoa(i) + ".html"
	}
	Crawl(url)
	page <- i
} //同时爬取不同页面

func Run(start, end int) {
	fmt.Printf("正在爬取第%d页到%d页\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go CrawlPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page) //防止爬虫还没有结束函数就退出
	}
}
