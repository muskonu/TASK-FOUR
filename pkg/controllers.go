package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

var db *gorm.DB

func init() {
	db = CreateDb()
}

func GetBook(context *gin.Context) {
	var book Book
	rand.Seed(time.Now().UnixNano())
	number := ReadNumber()
	id := rand.Intn(number) + 1
	db.First(&book, id)
	context.JSON(200, &book)
	fmt.Println(context.ContentType())
}

func GetBookByName(context *gin.Context) {
	var name string
	var book Book
	name = context.Param("name")
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
