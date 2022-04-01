package controllers

import (
	u "app/pdfGenerator"
	"app/resources"
	res "app/responses"
	"app/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func GetActiveProfileInfo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	profile, err := services.GetActiveProfileInfo(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_RETRIEVED",
		Data:    resources.ProfileResource(profile),
	})
}

func GetActiveProfile(w http.ResponseWriter, r *http.Request) {

	profile, err := services.GetActiveProfile()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_RETRIEVED",
		Data:    resources.ActiveProfileResource(profile),
	})
}

func GetActiveProfileSkills(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	skills, err := services.GetActiveProfileSkills(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_SKILLS",
		Data:    resources.SkillsResources(skills),
	})
}

func GetActiveProfilePortfolio(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	portfolio, err := services.GetActiveProfilePortfolio(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_PORTFOLIO",
		Data:    resources.PortfoliosResources(portfolio),
	})
}

func GetActiveProfileExperiences(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	experiences, err := services.GetActiveProfileExperiences(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_EXPERIENCES",
		Data:    resources.ExperiencesResources(experiences),
	})
}

func GetActiveProfileEducations(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	educations, err := services.GetActiveProfileEducations(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_EDUCATIONS",
		Data:    resources.EducationsResources(educations),
	})
}

func CreatePdf(w http.ResponseWriter, r *http.Request) {

	pdf := u.NewRequestPdf("")

	//html template path
	templatePath := "templates/template.html"

	//path for download pdf
	outputPath := "storage/example.pdf"

	//html template data
	templateData := struct {
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Jhon Lewis",
		Contact:     "Maria Anders",
		Country:     "Germany",
	}

	if err := pdf.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := pdf.GeneratePDF(outputPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println("ERROR: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templateData)
}
