package helpers

import (
	"app/config"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RequestRoute(r *http.Request, key ...string) (map[string]int, error) {

	result := make(map[string]int)

	for _, v := range key {
		key := mux.Vars(r)[v]
		if len(key) > 0 {
			value, err := strconv.Atoi(key)
			if err != nil {
				return result, errors.New("cannot parse route parameters")
			}
			result[v] = value
		}
	}

	return result, nil
}

func RequestRoute2(r *http.Request, s string) string {

	return mux.Vars(r)[s]
}

func InArray(x int, arr []int) bool {
	for _, v := range arr {
		if x == v {
			return true
		}
	}

	return false
}

type RouteParameters struct {
	UserId       int `route:"user_id"`
	ProfileId    int `route:"profile_id"`
	SkillId      int `route:"skill_id"`
	PortfolioId  int `route:"portfolio_id"`
	ExperienceId int `route:"experience_id"`
	EducationId  int `route:"education_id"`
}

func RouteParams(r *http.Request) RouteParameters {
	var routeParams RouteParameters

	enc := config.Encode{
		Struct: &routeParams,
		GetValue: func(s string) string {
			return mux.Vars(r)[s]
		},
	}

	enc.EnvEncode2("route")

	return routeParams
}
