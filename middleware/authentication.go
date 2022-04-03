package middleware

import (
	res "app/responses"
	"app/system"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Authentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !system.Verify(r) {
			res.JsonResponse(&w, res.UnauthorizedResponse{Message: "User is not authorized"})
			return
		}

		userId, err := strconv.Atoi(mux.Vars(r)["user_id"])
		if err != nil {
			fmt.Println("Middleware error = ", err.Error())
			res.JsonResponse(&w, res.BadRequestResponse{Message: "something went wrong"})
			return
		}

		if userId != system.Auth.UserId {
			res.JsonResponse(&w, res.UnauthorizedResponse{Message: "User is not authorized"})
			return
		}

		h.ServeHTTP(w, r)
	})
}
