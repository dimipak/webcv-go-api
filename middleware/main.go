package middleware

import (
	rw "app/responses"
	"app/systemService/authentication"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Authentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		res := rw.ResponseWriter{W: &w}
		if !authentication.Verify(r) {
			res.Unauthorized()
			return
		}

		userId, err := strconv.Atoi(mux.Vars(r)["user_id"])
		if err != nil {
			fmt.Println("Middleware error = ", err.Error())
			res.BadRequest("something went wrong")
			return
		}

		if userId != authentication.Auth.UserId {
			res.Unauthorized()
			return
		}

		h.ServeHTTP(w, r)
	})
}
