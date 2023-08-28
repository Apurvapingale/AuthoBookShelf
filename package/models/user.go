package models

import "github.com/Apurvapingale/book-store/package/config"

func GetUserData(userId int) User {
	db := config.GetDB()
	var user User
	db.Model(User{}).Where("id=?", userId).Find(&user)
	return user
}

func InActive(userId int) User {
	db := config.GetDB()
	var user User
	db.Model(User{}).Where("id=?", userId).Update("status", "INACTIVE")
	return user
}

func DeleteUser(userId int) User {
	db := config.GetDB()
	var user User
	db.Model(User{}).Where("id=?", userId).Update("status", "DELETED")
	return user
}

func AddReview(review RatingReview) RatingReview {
	db := config.GetDB()
	db.Model(RatingReview{}).Create(&review)
	return review
}
