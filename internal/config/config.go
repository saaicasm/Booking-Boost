package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// App Config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Errorlog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
