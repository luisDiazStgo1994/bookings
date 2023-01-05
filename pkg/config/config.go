package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// holds the app config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
