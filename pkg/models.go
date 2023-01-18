package pkg

import (
	"time"
)

type Book struct {
	BookId int `json:"book_id" gorm:"AUTO_INCREMENT;primary_key" form:"id"`
	Number int `json:"number" `
	SeBook
}

type SeBook struct {
	Author         string `json:"author" selector:"div.bd > h3.list-title > em" form:"author"`
	Name           string `json:"name" selector:"div.bd > h3.list-title > a" form:"name"`
	Classification string `json:"classification" selector:"div.list-content > p > a" form:"class"`
}

type Member struct {
	MemberId    int    `json:"member_id"`
	MemberLevel int    `json:"member_level"`
	HoldingBook int    `json:"holding_book"`
	Account     string `json:"account"`
	Password    string `json:"password"`
}

type Borrow struct {
	MemberId         int       `json:"member_id" form:"member_id"`
	BookId           int       `json:"book_id" form:"book_id"`
	BorrowDate       time.Time `json:"borrow_date" time_format:"2006-01-02 15:04:05" form:"borrow_date"`
	ExpectReturnDate time.Time `json:"expect_return_date" time_format:"2006-01-02 15:04:05" form:"expect_return_date"`
}

type Return struct {
	MemberId   int       `json:"member_id" form:"member_id"`
	BookId     int       `json:"book_id" form:"book_id"`
	BorrowDate time.Time `json:"borrow_date" time_format:"2006-01-02 15:04:05" form:"borrow_date"`
	ReturnDate time.Time `json:"return_date" time_format:"2006-01-02 15:04:05" form:"return_date"`
}

type LoginInfo struct {
	Account  string `json:"account" form:"account" binding:"required,number,min=9,max=15"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=15"`
}
