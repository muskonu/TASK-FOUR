package main

import (
	"ggggggggggggo/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/book", pkg.GetBook)
	r.GET("/books/:class/:number", pkg.GetBookByClass)
	r.GET("/books/:class", pkg.GetBookByClass)
	r.GET("/book/:name", pkg.GetBookByName)
	r.Run(":8080")
}
