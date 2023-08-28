package models

import (
	"github.com/Apurvapingale/book-store/package/config"
	"gorm.io/gorm"
)

// based on model which gives structure to the table
type Book struct {
	gorm.Model
	Name      string  `json:"name"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`                         //for user value is user and admin will have admin value
	Status   string `json:"status" gorm:"default:ACTIVE"` //for user deafault value is active

}

func (b *Book) CreateBook() *Book {
	db := config.GetDB()
	db.Model(Book{}).Create(&b)
	return b
}

func GetAllBooks() []Book {
	db := config.GetDB()
	var Books []Book
	db.Model(Book{}).Find(&Books)
	return Books
}

func GetBookById(Id int) *Book {
	db := config.GetDB()
	var getBook Book
	db.Model(Book{}).Where("id=?", Id).First(&getBook)
	return &getBook
}

func DeleteBook(Id int) Book {
	db := config.GetDB()
	var book Book
	db.Model(Book{}).Where("id?", Id).Delete(book)
	return book
}

func (b *User) RegisterUser() *User {
	db := config.GetDB()
	db.Model(User{}).Create(&b)
	return b
}
func (b *User) RegisterAdmin() *User {
	db := config.GetDB()
	db.Model(User{}).Create(&b)
	return b
}
