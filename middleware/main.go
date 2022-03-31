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
			res.JsonResponse(&w, res.UnauthorizedResponse{W: &w})
			return
		}

		userId, err := strconv.Atoi(mux.Vars(r)["user_id"])
		if err != nil {
			fmt.Println("Middleware error = ", err.Error())
			res.JsonResponse(&w, res.BadRequestResponse{
				Message: "something went wrong",
			})
			return
		}

		if userId != authentication.Auth.UserId {
			res.JsonResponse(&w, res.UnauthorizedResponse{W: &w})
			return
		}

		h.ServeHTTP(w, r)
	})
}

// func RouteParams(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		routeParams := config.Encode{
// 			Struct: &config.G_ROUTE_PARAMS,
// 			GetValue: func(s string) string {
// 				return mux.Vars(r)[s]
// 			},
// 		}

// 		routeParams.EnvEncode2("route")

// 		h.ServeHTTP(w, r)
// 	})
// }
