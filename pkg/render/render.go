package render

import (
	"bytes"
	"fmt"
	"github.com/KyawKyawThar/gowebtest/pkg/config"
	"github.com/KyawKyawThar/gowebtest/pkg/model"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var funcMap = template.FuncMap{}

var app *config.AppConfig

//New Templates set the config for the template package

func NewTemplates(a *config.AppConfig) {
	app = a
}

func DefaultTemplate(td *model.TemplateData) *model.TemplateData {
	return td
}

func Template(w http.ResponseWriter, tmp string, data *model.TemplateData) {
	//fmt.Println("HEHE", app.TemplateCache)
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	fmt.Println("Render tc", tc)

	t, ok := tc[tmp]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	data = DefaultTemplate(data)

	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

//Create a template cache as a map DS like about.page--> home.page

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {

		return myCache, err

	}

	for _, page := range pages {
		name := filepath.Base(page)

		//fmt.Println("pages", page)

		ts, err := template.New(name).Funcs(funcMap).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		fmt.Println("base", matches)
		if err != nil {
			return myCache, err

		}

		if len(matches) > 0 {

			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")

			if err != nil {
				return myCache, err
			}

		}

		myCache[name] = ts

	}

	return myCache, err
}
