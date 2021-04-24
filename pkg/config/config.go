package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

//AppConfig sets the app level configuration properties
type AppConfig struct {
	UseCache     bool
	TemplateCach map[string]*template.Template
	InProduction bool
	Session      *scs.SessionManager
}
