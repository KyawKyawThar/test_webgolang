package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

//AppConfig hold the application config

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Infolog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
