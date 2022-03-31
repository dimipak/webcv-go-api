package routes

// var G_ROUTE_PARAMS RouteParams

// type RouteParams struct {
// 	UserId    int `route:"user_id"`
// 	ProfileId int `route:"profile_id"`
// 	SkillId   int `route:"skill_id"`
// }

// func RouteParams(r *http.Request) config.RouteParams {
// 	var routeParams config.RouteParams

// 	enc := config.Encode{
// 		Struct: &routeParams,
// 		GetValue: func(s string) string {
// 			return mux.Vars(r)[s]
// 		},
// 	}

// 	enc.EnvEncode2("route")

// 	return routeParams
// }
