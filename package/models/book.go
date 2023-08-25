package models

import (
	"github.com/Apurvapingale/book-store/package/config"

	"gorm.io/gorm"
)

var db  *gorm.DB

//based on model which gives structure to the table
type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
}


type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`//for user value is user and admin will have admin value
	Status string `json:"status" gorm:"default:ACTIVE"`//for user deafault value is active 

}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID?", Id).Delete(book)
	return book
}

func (b *User) RegisterUser() *User {
	db.Create(&b)
	return b
}
