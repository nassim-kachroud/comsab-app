package healthz

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/database"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/logger"
)

//Controller ...
type Controller struct {
	Manager HealthzManager
}

//New ...
func New(manager HealthzManager) Controller {
	return Controller{Manager: manager}
}

//Healthz ...
func (c Controller) Healthz(w http.ResponseWriter, r *http.Request) {
	healthStatus := HealthStatus{}
	statusCode := http.StatusOK
	healthStatus.Databases.Ready = true

	//Ping DbNew
	healthStatus.Databases.DbNew = c.Manager.getDatabaseStatus(database.DB{Pinger: database.DbNew})

	//Set status code according to errors
	if len(healthStatus.Databases.DbCompta.Status) == 0 {
		statusCode = http.StatusInternalServerError
	}

	healthStatus.Status = statusCode

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	data, err := json.MarshalIndent(&healthStatus, "", "  ")

	if err != nil {
		logger.Log.Error(errors.Wrap(err, "Healthz Status : json.MarshalIndent").Error())
	}
	w.Write(data)
}
