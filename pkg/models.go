package pkg

type Book struct {
	Id uint `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	SeBook
}

type SeBook struct {
	Author         string `json:"author" gorm:"not null;check:" selector:"div.bd > h3.list-title > em"`
	Name           string `json:"name" gorm:"not null" selector:"div.bd > h3.list-title > a"`
	Classification string `json:"classification" gorm:"not null" selector:"div.list-content > p > a"`
}
