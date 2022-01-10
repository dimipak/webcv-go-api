package controllers

import (
	u "app/pdfGenerator"
	"app/resources"
	"app/responses"
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

	res := responses.ResponseWriter{W: &w}
	if err != nil {
		res.BadRequest("Something happened")
		return
	}

	res.Success("PROFILE_RETRIEVED", resources.ProfileResource(profile))
}

func GetActiveProfile(w http.ResponseWriter, r *http.Request) {

	profile, err := services.GetActiveProfile()

	res := responses.ResponseWriter{W: &w}
	if err != nil {
		res.BadRequest("Something happened")
		return
	}

	res.Success("ACTIVE_PROFILE_RETRIEVED", resources.ActiveProfileResource(profile))
}

func GetActiveProfileSkills(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	skills, err := services.GetActiveProfileSkills(profileId)

	res := responses.ResponseWriter{W: &w}
	if err != nil {
		res.BadRequest("Something happened")
		return
	}

	res.Success("ACTIVE_PROFILE_SKILLS", resources.SkillResources(skills))
}

func GetActiveProfilePortfolio(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	portfolio, err := services.GetActiveProfilePortfolio(profileId)

	res := responses.ResponseWriter{W: &w}

	if err != nil {
		res.BadRequest("Something happened")
		return
	}

	res.Success("ACTIVE_PROFILE_PORTFOLIO", resources.PortfolioResources(portfolio))
}

func GetActiveProfileExperiences(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	experiences, err := services.GetActiveProfileExperiences(profileId)

	res := responses.ResponseWriter{W: &w}

	if err != nil {
		res.BadRequest("Something happened")
		return
	}

	res.Success("ACTIVE_PROFILE_EXPERIENCES", resources.ExperienceResources(experiences))
}

func GetActiveProfileEducations(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	educations, err := services.GetActiveProfileEducations(profileId)

	res := responses.ResponseWriter{W: &w}

	if err != nil {
		res.BadRequest("Something happened")
		return
	}

	res.Success("ACTIVE_PROFILE_EDUCATIONS", resources.EducationResources(educations))
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
