package responses

type SuccessResponse struct {
	Message string
	Data    interface{}
}

type BadRequestResponse struct {
	Message string
}

type UnauthorizedResponse struct {
	Message string
}

type SuccessWithTokenResponse struct {
	Message string
	Token   string
	Data    interface{}
}

type MethodNotAllowedResponse struct {
	Message string
}

type NotFoundResponse struct {
	Message string
}
