package middleware

import (
	"fmt"
	"net/http"

	"github.com/Apurvapingale/book-store/package/auth"
	"github.com/Apurvapingale/book-store/package/helper"
)

func ValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "request does not contain an access token", 401)
			return
		}

		tokenString, tokenerror := helper.StripBearerPrefixFromTokenString(tokenString)

		if tokenerror != nil {
			http.Error(w, "error while parsing the authorization token", http.StatusGone)

			return
		}
		fmt.Println("got jwt ", tokenString)
		err := auth.ValidateToken(tokenString, "USER")
		if err != nil {

			http.Error(w, err.Error(), http.StatusGone)

			return
		}
		next.ServeHTTP(w, r)
	})

}

func ValidateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "request does not contain an access token", 401)
			return
		}

		tokenString, tokenerror := helper.StripBearerPrefixFromTokenString(tokenString)

		if tokenerror != nil {
			http.Error(w, "error while parsing the authorization token", http.StatusGone)

			return
		}
		fmt.Println("got jwt ", tokenString)
		err := auth.ValidateToken(tokenString, "ADMIN")
		if err != nil {

			http.Error(w, err.Error(), http.StatusGone)

			return
		}
		next.ServeHTTP(w, r)
	})

}
