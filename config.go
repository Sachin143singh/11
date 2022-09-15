// package config

// import (
// 	"html/template"
// 	"log"
// )

// // AppConfig holds the application config
// type AppConfig struct {
// 	UseCache      bool
// 	TemplateCache map[string]*template.Template
// 	InfoLog       *log.Logger
// }

package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
