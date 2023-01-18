package main

import (
	"ggggggggggggo/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	go pkg.CheckTiker()
	pkg.LogFile()
	r := gin.Default()
	r.GET("/book", pkg.GetBook)
	r.GET("/create", pkg.Auth, pkg.CreateBook)
	r.GET("/borrow", pkg.Auth, pkg.BorrowBook)
	r.GET("/return", pkg.Auth, pkg.ReturnBook)
	r.GET("/login", pkg.Login)
	r.GET("/register", pkg.Register)
	r.GET("info", pkg.Auth, pkg.Info)
	r.Run(":8080")
}
