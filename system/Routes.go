package system

import (
	"app/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

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

	decoder := helpers.Decoder{
		Interface: &routeParams,
		GetValue: func(s string) string {
			return mux.Vars(r)[s]
		},
	}

	decoder.Decode("route")

	return routeParams
}
