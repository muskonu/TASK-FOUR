package pkg

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func GetBook(context *gin.Context) {
	var book Book
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(10) + 1
	db := CreateDb()
	db.First(&book, id)
	context.JSON(200, &book)
}

func GetAllbook(context *gin.Context) {
	var books []Book
	db := CreateDb()
	db.Find(&books)
	context.JSON(200, books)
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

}
