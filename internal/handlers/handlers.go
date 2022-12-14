package handlers

import (
	"net/http"

	"github.com/Ahdeyyy/go-web/internal/config"
)

var Dep *Dependency

// Dependency is the dependency for the handlers
type Dependency struct {
	// App is the application configuration
	App *config.Config
}

// NewDependency creates a new dependency for the handlers
func NewDependency(app *config.Config) *Dependency {
	return &Dependency{
		App: app,
	}
}

// Init initializes the handlers
func Init(d *Dependency) {
	Dep = d
}

// Home is the home handler
func (d *Dependency) Home(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	w.Write([]byte("Hello World"))
}
