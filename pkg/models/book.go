package models

import (
	"readinglist-api/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model        // using gorm convention for auto id, created_at, updated_at, deleted_at
	Title      string `json:"title"`
	Author     string `json:"author"`
	Category   string `json:"category"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) AddBook() *Book {
    db.Create(&b)
    return b
}

func ViewAll() []Book {
    var Books []Book
    db.Find(&Books)
    return Books
}

func GetBook(Id int64) (*Book, *gorm.DB) {
    var getBook Book
    db := db.Where("ID=?", Id).Find(&getBook)
    return &getBook, db
}

func Delete(Id int64) Book {
    var book Book
    db.Where("ID=?", Id).Delete(book)
    return book
}

