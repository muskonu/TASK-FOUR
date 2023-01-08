package pkg

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

func GetBook(context *gin.Context) {
	var book Book
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(160) + 32
	db := CreateDb()
	db.First(&book, id)
	context.JSON(200, &book)
}

func GetBookByName(context *gin.Context) {
	var name string
	var book Book
	name = context.Param("name")
	db := CreateDb()
	db.Where("name=?", name).First(&book)
	if book == (Book{}) {
		context.String(200, "没有这本书呢亲")
	} else {
		context.JSON(200, &book)
	}
}

func GetBookByClass(context *gin.Context) {
	var class string
	var books []Book
	var number string
	var n int
	class = context.Param("class")
	db := CreateDb()
	number = context.Param("number")
	n, _ = strconv.Atoi(number)
	if n == 0 {
		n = 1
	}
	db.Where("classification=?", class).Limit(n).Find(&books)
	if Judge(class) {
		context.String(200, "没有这类书呢亲")
	} else {
		context.JSON(200, &books)
	}
}
