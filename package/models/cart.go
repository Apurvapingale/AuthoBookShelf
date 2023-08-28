package models

import (
	"github.com/Apurvapingale/book-store/package/config"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	BookId    int     `json:"book_id"`
	UserId    int     `json:"user_id"`
	BookName  string  `json:"book_name"`
	BookPrice float32 `json:"book_price"`
	Quantity  int     `json:"quantity"`
	AmtTotal  float32 `json:"amt_total"` //it is multiplication of book price and quantity
}

type Order struct {
	gorm.Model
	OrderId       string  `json:"order_id"`
	UserId        int     `json:"user_id"`
	Quantity      int     `json:"quantity"`
	Total         float32 `json:"total"` //exact total amount payed by user
	PaymentStatus string  `json:"payment_status"`
}

type OrderDetail struct {
	gorm.Model
	OrderId   string  `json:"order_id"`
	BookId    int     `json:"book_id"`
	BookName  string  `json:"book_name"`
	BookPrice float32 `json:"book_price"`
	Quantity  int     `json:"quantity"`
	AmtTotal  float32 `json:"amt_total"` //it is multiplication of book price and quantity
}

type RatingReview struct {
	gorm.Model
	UserId int     `json:"user_id"`
	BookId int     `json:"book_id"`
	Rating float32 `json:"rating"`
	Review string  `json:"review"`
}

func GetOrdersByUserID(userId int) []Order {
	db := config.GetDB()
	var orders []Order
	db.Model(Order{}).Where("user_id=?", userId).Find(&orders)
	return orders
}
