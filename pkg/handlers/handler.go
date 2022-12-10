package handlers

import (
	"fmt"
	"github.com/KyawKyawThar/gowebtest/pkg/config"
	"github.com/KyawKyawThar/gowebtest/pkg/model"
	"github.com/KyawKyawThar/gowebtest/pkg/render"
	"net/http"
)

//NOTE Repository pattern allow us to swap components
// within a minimum changes required code base.....

// Repo the repository use by handler (User in Main.go)
var Repo *Repository

// Repository is the type of repo
type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	//data come from main (a) and save in Repository struct.because we need to use that data
	return &Repository{App: a}
}

// RepoHandler sets handler for repo
func RepoHandler(r *Repository) {
	Repo = r

	fmt.Println("Repo", *Repo)
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	//add business login in model.TemplateData
	render.Template(w, "home.page.gohtml", &model.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	mapString := map[string]string{}

	mapString["Hey Gopher"] = "Blockchain is my dream"

	//add business login in model.TemplateData

	render.Template(w, "about.page.gohtml", &model.TemplateData{
		StringMap: mapString,
	})
}
