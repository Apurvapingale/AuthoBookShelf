package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Apurvapingale/book-store/package/auth"
	"github.com/Apurvapingale/book-store/package/config"
	"github.com/Apurvapingale/book-store/package/models"
	"github.com/Apurvapingale/book-store/package/utils"
	"github.com/gorilla/mux"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := utils.ParseBody(r, &newUser)
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error while parsing"}`))
		return
	}
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error while hashing password"}`))
		return
	}
	newUser.Password = string(hashedpassword)

	newUser.Status = "ACTIVE"
	newUser.Role = "USER"
	resp := newUser.RegisterUser()
	respBytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error converting response to bytes ", err.Error()) //this to console
		w.Header().Set("Content-Type", "pkglication/json")              //this to client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "log the error and send response to client"}`))
		return
	} //log the error and send response to client
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return

}

func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := utils.ParseBody(r, &newUser)
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error while parsing"}`))
		return
	}
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error while hashing password"}`))
		return

	}
	newUser.Password = string(hashedpassword)

	newUser.Status = "ACTIVE"
	newUser.Role = "ADMIN"
	resp := newUser.RegisterUser()
	respBytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error converting response to bytes ", err.Error()) //this to console
		w.Header().Set("Content-Type", "pkglication/json")              //this to client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "log the error and send response to client"}`))
		return
	} //log the error and send response to client
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return

}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := utils.ParseBody(r, &user)
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error while parsing"}`))
		return
	}
	var userDBdata models.User
	db := config.GetDB()
	db.Where("email=?", user.Email).First(&userDBdata)
	if userDBdata.ID == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "user not found"}`))
		return
	}
	if userDBdata.Status != "ACTIVE" {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "user is not active"}`))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userDBdata.Password), []byte(user.Password))
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "invalid password"}`))
		return
	}
	token, err := auth.GenerateJWT(userDBdata.ID, userDBdata.Role, userDBdata.Email)
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error while generating token"}`))
		return
	}

	//log the error and send response to client
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	var resp = map[string]interface{}{"token": token}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error converting response to bytes ", err.Error()) //this to console
		w.Header().Set("Content-Type", "pkglication/json")              //this to client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "log the error and send response to client"}`))
		return
	}

	w.Write(respBytes)
	return

}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Error while parsing")                 //this to console
		w.Header().Set("Content-Type", "pkglication/json") //this to client
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Unable to parse the request"}`))
	}
	userDetails := models.GetUserData(ID)

	//send response to user
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

func InActive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Error while parsing")                 //this to console
		w.Header().Set("Content-Type", "pkglication/json") //this to client
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Unable to parse the request"}`))
	}
	userDetails := models.InActive(ID)

	//send response to user
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Error while parsing")                 //this to console
		w.Header().Set("Content-Type", "pkglication/json") //this to client
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Unable to parse the request"}`))
	}
	userDetails := models.DeleteUser(ID)

	//send response to user
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

func AddReview(w http.ResponseWriter, r *http.Request) {
	var newReview models.RatingReview
	err := utils.ParseBody(r, &newReview)
	if err != nil {
		w.Header().Set("Content-Type", "pkglication/json") //this to client
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error while parsing"}`))
		return
	}
	review := models.AddReview(newReview)

	//send response to user
	res, _ := json.Marshal(review)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}
