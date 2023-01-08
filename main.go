package main

import (
	"ggggggggggggo/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Run(1, 5)
	r.GET("/book", pkg.GetBook)
	r.GET("/books", pkg.GetAllbook)
	r.GET("/book/:name", pkg.GetBookByName)
	r.Run(":8080")
}
