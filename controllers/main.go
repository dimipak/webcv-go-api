package controllers

import (
	"app/requests"
	"app/services"
)

// Request
var (
	userRegisterRequest     requests.UserRegisterRequest
	updateProfileRequest    requests.UpdateProfileRequest
	portfolioCreateRequest  requests.PortfolioCreateRequest
	portfolioUpdateRequest  requests.PortfolioUpdateRequest
	experienceCreateRequest requests.ExperienceCreateRequest
	experienceUpdateRequest requests.ExperienceUpdateRequest
	educationCreateRequest  requests.EducationCreateRequest
	educationUpdateRequest  requests.EducationUpdateRequest
	userLoginRequest        requests.UserLoginRequest
)

// Services
var (
	profileService    services.ProfileService
	portfolioService  services.PortfolioService
	experienceService services.ExperienceService
	educationService  services.EducationService
)
