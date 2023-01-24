package pkg

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

var db *gorm.DB

func init() {
	db = CreateDb()
}

func GetBook(context *gin.Context) { //查询书籍信息
	var book Book
	err := context.ShouldBindQuery(&book)
	if err != nil {
		context.String(200, "查询错误")
	}
	result := db.Where(&book).Order("rand()").First(&book)
	if result.Error != nil {
		context.String(200, "查询错误")
	} else {
		context.JSON(200, &book)
	}
}

func CreateBook(context *gin.Context) { //存入书籍
	var book Book
	var member Member
	err := context.ShouldBindQuery(&book)
	if err != nil {
		context.JSON(400, gin.H{"err": "创造出错"})
	}
	if book.Name == "" || book.Author == "" || book.Classification == "" {
		context.JSON(401, gin.H{"err": "请补全信息"})
		return
	}
	if !Judge(book.Classification) {
		context.JSON(400, gin.H{"err": "没有这类分类"})
		return
	}
	result := db.Where(&book).First(&book)
	memberId, _ := context.Get("member_id") //上交一本书，等级加一,最高为4，可持有书本总数和持有天数增加
	id, _ := strconv.Atoi(memberId.(string))
	db.Where(&Member{MemberId: id}).First(&member)
	member.MemberLevel += 1
	db.Save(&member)
	if result.Error != nil {
		db.Save(&book)
		context.JSON(200, gin.H{"result": "入库成功"})
	} else {
		book.Number += 1
		db.Save(&book)
		context.JSON(200, gin.H{"result": "入库成功"})
	}
}

func BorrowBook(context *gin.Context) {
	var book Book
	var borrows []Borrow
	id := context.Query("id") //查找书籍
	if len(id) == 0 {         //若不输入书籍id则返回所有借阅记录
		memberId, _ := context.Get("member_id")
		i, _ := strconv.Atoi(memberId.(string))
		db.Where(&Borrow{MemberId: i}).Find(&borrows)
		context.JSON(200, borrows)
		return
	}
	book.BookId, _ = strconv.Atoi(id) //若输入书籍id则进行判定,若信息符合规范则插入并返回所有借阅记录
	result := db.Where(&book).First(&book)
	if result.Error != nil {
		context.String(403, "没有此书")
		return
	}
	if book.Number == 0 {
		context.String(403, "此书以全部借出")
		return
	}
	memberId, _ := context.Get("member_id") //获得借阅者信息
	i, _ := strconv.Atoi(memberId.(string))
	db.Exec("call borrowbook(?,?)", i, book.BookId) //调用存储过程
	if err := db.Where(&Borrow{MemberId: i}).Find(&borrows).Error; err != nil {
		context.JSON(403, gin.H{"err": err})
		return
	}
	context.JSON(200, borrows)
}

func Login(context *gin.Context) { //登录
	var loginInfo LoginInfo
	var member Member
	err := context.ShouldBind(&loginInfo)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	member.Account = loginInfo.Account
	member.Password = loginInfo.Password
	result := db.Where(&member).First(&member)
	if result.Error != nil {
		context.JSON(401, gin.H{"err": "账号或密码错误"})
		return
	}
	token, err := Token(member)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	context.SetCookie("token", token, 0, "/", "", false, true)
	context.JSON(200, gin.H{"result": "登陆成功"})
}

func Auth(context *gin.Context) { //验证token
	auth, err := context.Cookie("token")
	if err != nil {
		context.Abort()
		context.JSON(401, gin.H{
			"err": "请登录后尝试",
		})
		return
	}
	auth = strings.Fields(auth)[1]
	token, err := jwt.ParseWithClaims(auth, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("abaaba"), nil
	})
	if err != nil {
		context.Abort()
		context.JSON(401, gin.H{
			"err": "访问出现错误",
		})
		return
	}
	if claim, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		context.Set("member_id", claim.Id) //传递验证的成员id
	} else {
		context.Abort()
		context.JSON(401, gin.H{
			"err": "登陆验证已过期",
		})
		return
	}
}

func Register(context *gin.Context) {
	var loginInfo LoginInfo
	var member Member
	var pps [16]byte
	var secret = "salt"
	err := context.ShouldBind(&loginInfo)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	member.Account = loginInfo.Account
	member.Password = loginInfo.Password + secret
	pps = md5.Sum([]byte(member.Password))
	member.Password =string(pps[:16])
	err = db.Where(&member).First(&Member{}).Error
	fmt.Println(err)
	if err == nil {
		context.JSON(401, gin.H{"err": "该账号已被注册"})
		return
	} else {
		db.Create(&member)
		context.JSON(200, gin.H{"result": "您已注册成功"})
	}
}

func ReturnBook(context *gin.Context) {
	var returns []Return
	id := context.Query("id") //查找书籍
	memberId, _ := context.Get("member_id")
	i, _ := strconv.Atoi(memberId.(string))
	bid, _ := strconv.Atoi(id)
	if len(id) == 0 { //若不输入书籍id则返回所有借阅记录
		db.Where(&Return{MemberId: i}).Find(&returns)
		context.JSON(200, returns)
		return
	}
	db.Exec("call returnbook(?,?)", i, bid)
	if err := db.Where(&Borrow{MemberId: i}).Find(&returns).Error; err != nil {
		context.JSON(403, gin.H{"err": err})
		return
	}
	context.JSON(200, returns)
}

func Info(context *gin.Context) {
	var member Member
	memberId, _ := context.Get("member_id")
	id, _ := strconv.Atoi(memberId.(string))
	db.Where(&Member{MemberId: id}).Find(&member)
	context.JSON(200, &member)
}
