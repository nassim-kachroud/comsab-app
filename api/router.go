package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/database"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/healthz"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/invoice"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/overview"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/transaction"
)

const welcomeMessage = "Welcome to the invoicing API!"

//NewRouter api router
func NewRouter(mainRouter *mux.Router) http.Handler {
	//Root route
	mainRouter.HandleFunc("/", index).Methods("GET")

	//Metrics route
	mainRouter.Handle("/metrics", prometheus.Handler()).Methods("GET")

	apiRouter := mainRouter.PathPrefix("/").Subrouter().StrictSlash(true)

	//Endpoints
	// overview.NewRouter(apiRouter, overview.New(overview.Manager{Repository: overview.Repository{DbNew: database.DbNew}}))
	// invoice.NewRouter(apiRouter, invoice.New(invoice.Manager{Repository: invoice.Repository{Invoices: database.Invoices}, Client: &http.Client{}}))
	// transaction.NewRouter(apiRouter, transaction.New(transaction.Manager{Repository: transaction.Repository{Invoices: database.Invoices}}))
	healthz.NewRouter(apiRouter, healthz.New(healthz.Manager{Client: &http.Client{}}))

	return apiRouter
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, welcomeMessage)
}
