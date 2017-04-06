package healthz

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter ...
func NewRouter(mainRouter *mux.Router, c Controller) http.Handler {

	healthzRouter := mainRouter.PathPrefix("/").Subrouter().StrictSlash(true)

	healthzRouter.HandleFunc("/healthz", c.Healthz)

	return healthzRouter
}
