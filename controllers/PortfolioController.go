package controllers

import (
	"app/mappers"
	"app/models"
	"app/resources"
	res "app/responses"
	"app/system"
	"net/http"
)

type PortfolioController struct{}

func (p *PortfolioController) Get(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_RETRIEVED",
		Data:    resources.PortfoliosResources(profile.GetPortfolios().Portfolio),
	})
}

func (p *PortfolioController) Show(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := profile.GetPortfolios().Portfolio.GetOne(routes.PortfolioId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_RETRIEVED",
		Data:    resources.PortfolioResources(portfolio),
	})
}

func (p *PortfolioController) Create(w http.ResponseWriter, r *http.Request) {

	err := portfolioCreateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)

	_, err = profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := portfolioService.SetProfileId(routes.ProfileId).SetPortfolioCreateRequest(portfolioCreateRequest).Create(mappers.PortfolioCreateMapper)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_CREATED",
		Data:    resources.PortfolioResources(portfolio),
	})
}

func (p *PortfolioController) Update(w http.ResponseWriter, r *http.Request) {

	err := portfolioUpdateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := profile.GetPortfolios().Portfolio.GetOne(routes.PortfolioId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio.Update(models.Portfolio{
		Name:       portfolioUpdateRequest.Name,
		Type:       portfolioUpdateRequest.Type,
		Technology: portfolioUpdateRequest.Technology,
		Customer:   portfolioUpdateRequest.Customer,
		WebsiteUrl: portfolioUpdateRequest.WebsiteUrl,
	})

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_UPDATED",
		Data:    resources.PortfolioResources(portfolio),
	})
}

func (p *PortfolioController) UploadImage(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := profile.GetPortfolios().Portfolio.GetOne(routes.PortfolioId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	img := system.ImageUpload{
		Allowed:  "portfolio_image",
		Path:     "portfolio/",
		FileType: "jpg",
	}

	url, err := img.Upload(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio.Update(models.Portfolio{
		ImageUrl: url,
	})

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_IMAGE_UPDATEDD",
		Data:    resources.PortfolioResources(portfolio),
	})
}

func (p *PortfolioController) Delete(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := profile.GetPortfolios().Portfolio.GetOne(routes.PortfolioId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio.Delete()

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_DELETED",
		Data:    resources.PortfolioResources(portfolio),
	})
}
