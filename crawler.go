package main

import (
	"fmt"
	"ggggggggggggo/pkg"
	"github.com/gocolly/colly"
)

func Crawl(url string) {
	var books []pkg.Book
	var number int64
	c := colly.NewCollector()
	db := pkg.CreateDb()
	db.Find(&books).Count(&number)
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnHTML("#category a[href]", func(e *colly.HTMLElement) {
		if e.Attr("href") != "/dianzishuzhizuo/" {
			e.Request.Visit(e.Attr("href"))
		}
	})
	c.OnHTML("div.pagenavi a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnHTML("div.channel-item", func(e *colly.HTMLElement) {
		sbook := pkg.SeBook{}
		err := e.Unmarshal(&sbook)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		book := pkg.Book{}
		book.SeBook = sbook
		book.BookId = int(number + 1)
		err = db.Save(&book).Error
		if err == nil {
			number++
		}
	})
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Visit error:", err)
		return
	}
	pkg.WriteNumber(int(number))
} //爬取该页面的作者，书名以及分页
