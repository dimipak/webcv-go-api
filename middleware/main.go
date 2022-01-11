package middleware

import (
	res "app/responses"
	"app/systemService/authentication"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Authentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !authentication.Verify(r) {
			res.JsonResponse(res.UnauthorizedResponse{W: &w})
			return
		}

		userId, err := strconv.Atoi(mux.Vars(r)["user_id"])
		if err != nil {
			fmt.Println("Middleware error = ", err.Error())
			res.JsonResponse(res.BadRequestResponse{
				W:       &w,
				Message: "something went wrong",
			})
			return
		}

		if userId != authentication.Auth.UserId {
			res.JsonResponse(res.UnauthorizedResponse{W: &w})
			return
		}

		h.ServeHTTP(w, r)
	})
}
