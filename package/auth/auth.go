package auth

import (
	"errors"
	"log"
	"time"

	"github.com/Apurvapingale/book-store/package/config"
	"github.com/Apurvapingale/book-store/package/models"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("supersecretkeyvdjwbdhwjdbiwuhdqwihdiq")

type JWTClaim struct {
	Id     uint   `json:"id"`
	Role  string `json:"role"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(id uint, role string, email string) (tokenString string, err error) {
	//default exp time for jwt token is 24 hours 
	expirationTime := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	claims := &JWTClaim{
		Id:     id,
		Email:  email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "book-store",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

var UserJwtData *JWTClaim

func ValidateToken(signedToken string, role string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		log.Fatal("error occurred during parsing the token", err.Error())
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		log.Fatal("error occurred during parsing the token", err.Error())
		return
	}

	if token.Valid {
		UserJwtData = claims
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		//fmt.Println("That's not even a token")
		err = errors.New("that's not even a token")
		log.Println("That's not even a token", err.Error())
		return
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		//fmt.Println("Timing is everything")
		err = errors.New("token is either expired or not active yet")
		log.Println("token is either expired or not active yet", err.Error())
		return
	} else {
		//fmt.Println("Couldn't handle this token:", err)
		err = errors.New("couldn't handle this token")
		log.Println("Couldn't handle this token:", err.Error())
		return
	}
	//get user data from db and validite its status wheather it is active or not also validate the role of the user

    db := config.GetDB()
   var user models.User
   db.Where("id=?", claims.Id).Find(&user)
   if user.Status != "ACTIVE" {
	   err = errors.New("user is not active")
	   log.Println("user is not active", err.Error())
	   return
	      }
		     if user.Role != role {
				 err = errors.New("user is not authorized")
				 log.Println("user is not authorized", err.Error())
				 return
				}



	return
}
