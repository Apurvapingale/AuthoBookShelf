package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Apurvapingale/book-store/package/config"
	"github.com/Apurvapingale/book-store/package/models"
	"gorm.io/gorm"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {

	userid, _ := r.Context().Value("userId").(int)
	bookid, _ := r.Context().Value("bookId").(int)
	if userid == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "userid is required"}`))

		return
	}
	if bookid == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "bookid is required"}`))

		return
	}
	bookdata := models.GetBookById(bookid)
	if bookdata.ID == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "bookid is not valid"}`))

		return
	}
	if bookdata.Quantity <= 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "book is out of stock"}`))

		return
	}
	if bookdata.Quantity >= 1 {
		var cartData models.CartItem
		db := config.GetDB()
		er := db.Model(models.CartItem{}).Where("user_id = ? AND book_id = ?", userid, bookid).First(&cartData)
		if er.Error != nil {
			log.Println(er)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if cartData.ID == 0 {

			cartData.UserId = userid
			cartData.BookId = bookid
			cartData.BookName = bookdata.Name
			cartData.BookPrice = bookdata.Price
			cartData.Quantity = 1
			cartData.AmtTotal = bookdata.Price

		} else { //present in cart
			cartData.Quantity += 1
			cartData.AmtTotal = float32(cartData.Quantity) * cartData.BookPrice

		}
		err := db.Model(models.CartItem{}).Save(&cartData)
		if err.Error != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return

		}
		// Prepare the response message
		response := map[string]interface{}{
			"msg": "product added to cart",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func AddToOrder(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user ID from the context
	userid, _ := r.Context().Value("userId").(int)

	if userid == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "userid is required"}`))
		return
	}

	var cartData []models.CartItem

	db := config.GetDB()
	tx := db.Begin()

	er := tx.Model(models.CartItem{}).Where("user_id = ? ", userid).Find(&cartData)
	if er.Error != nil {
		log.Println(er.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if len(cartData) > 0 {
		// Create a new order entry
		orderId, err := generateUniqueID()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var totalAmt float32
		var totalQty int

		order := models.Order{
			OrderId:       orderId,
			UserId:        userid,
			Quantity:      totalQty,
			Total:         totalAmt,
			PaymentStatus: "PAID",
		}

		err2 := tx.Model(models.Order{}).Create(&order)
		if err2.Error != nil {
			log.Println(err2)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		for _, v := range cartData {

			//get book data and check if quantity is available
			temp_book := models.GetBookById(v.BookId)
			if temp_book.Quantity < v.Quantity {
				tx.Rollback()
				log.Println("book is out of stock")
				w.Header().Set("Content-Type", "pkglication/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte(`{"message": "book is out of stock"}`))

				return
			}

			amtTotal := float32(v.Quantity) * v.BookPrice
			orderDetail := models.OrderDetail{
				OrderId:   orderId,
				BookId:    v.BookId,
				BookName:  v.BookName,
				BookPrice: v.BookPrice,
				Quantity:  v.Quantity,
				AmtTotal:  amtTotal,
			}

			totalAmt += amtTotal
			totalQty += v.Quantity

			err := tx.Model(models.OrderDetail{}).Create(&orderDetail)
			if err.Error != nil {
				tx.Rollback()
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			err4 := DecrementBookQtyFromDb(tx, v.BookId, v.Quantity)
			if err4 != nil {
				tx.Rollback()
				log.Println(err4)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

		}
		err3 := tx.Model(models.Order{}).Where("order_id = ?", order.ID).Update("total", totalAmt).Update("quantity", totalQty)
		if err3.Error != nil {
			tx.Rollback()
			log.Println(err3)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err5 := RemoveCartFromDb(tx, userid)
		if err5 != nil {
			tx.Rollback()
			log.Println(err5)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tx.Commit()

		// Prepare the response message
		response := map[string]interface{}{
			"msg": "order sccessfully placed",
		}

		// Encode the response as JSON and write it to the response writer
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		// Product is already in the cart
		response := map[string]interface{}{
			"msg": "No items in cart",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func generateUniqueID() (string, error) {
	// Generate a random 6-byte ID
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes using base64 encoding
	uniqueID := base64.URLEncoding.EncodeToString(randomBytes)

	// Trim padding and non-alphanumeric characters
	for i := 0; i < len(uniqueID); i++ {
		if !((uniqueID[i] >= '0' && uniqueID[i] <= '9') || (uniqueID[i] >= 'a' && uniqueID[i] <= 'z') || (uniqueID[i] >= 'A' && uniqueID[i] <= 'Z')) {
			uniqueID = uniqueID[:i] + uniqueID[i+1:]
			i--
		}
	}

	// Take the first 8 characters as the final order ID
	if len(uniqueID) >= 8 {
		uniqueID = uniqueID[:8]
	} else {
		return "", fmt.Errorf("could not generate a unique ID")
	}

	return uniqueID, nil
}

func DecrementCart(w http.ResponseWriter, r *http.Request) {
	userid, _ := r.Context().Value("userId").(int)
	bookid, _ := r.Context().Value("bookId").(int)
	if userid == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "userid is required"}`))

		return
	}
	if bookid == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "bookid is required"}`))

		return
	}
	bookdata := models.GetBookById(bookid)
	if bookdata.ID == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "bookid is not valid"}`))

		return
	}

	var cartData models.CartItem
	db := config.GetDB()
	er := db.Model(models.CartItem{}).Where("user_id = ? AND book_id = ?", userid, bookid).First(&cartData)
	if er.Error != nil {
		log.Println(er)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if cartData.ID == 0 {

		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "book is not in cart"}`))
		return

	} else { //present in cart

		Qty := cartData.Quantity - 1
		if Qty == 0 {
			err := db.Model(models.CartItem{}).Where("user_id = ? AND book_id = ?", userid, bookid).Delete(&models.CartItem{}).Error
			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			// Prepare the response message
			response := map[string]interface{}{
				"msg": "Item removed from cart",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			w.WriteHeader(http.StatusOK)
			return
		}
		cartData.Quantity = Qty
		cartData.AmtTotal = float32(cartData.Quantity) * cartData.BookPrice

	}
	err := db.Model(models.CartItem{}).Save(&cartData)
	if err.Error != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return

	}
	// Prepare the response message
	response := map[string]interface{}{
		"msg": "Item Quantity decremented",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}

func RemoveCartFromDb(db *gorm.DB, userId int) (err error) {

	// Remove the cart items from the database for the user if the order is placed
	err = db.Model(models.CartItem{}).Where("user_id = ?", userId).Delete(&models.CartItem{}).Error
	return

}

func RemoveAllItemsFromCart(w http.ResponseWriter, r *http.Request) {
	userid, _ := r.Context().Value("userId").(int)
	if userid == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "userid is required"}`))

		return
	}
	db := config.GetDB()
	err := RemoveCartFromDb(db, userid)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Prepare the response message
	response := map[string]interface{}{
		"msg": "All Items removed from cart",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}

func DecrementBookQtyFromDb(db *gorm.DB, bookId int, qty int) (err error) {

	// Decrement the book quantity from the database
	err = db.Model(models.Book{}).Where("id = ?", bookId).Update("quantity", gorm.Expr("quantity - ?", qty)).Error
	return

}

func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	userid, _ := r.Context().Value("userId").(int)
	if userid == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "userid is required"}`))

		return
	}
	orders := models.GetOrdersByUserID(userid)
	resp, _ := json.Marshal(orders)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return
}
