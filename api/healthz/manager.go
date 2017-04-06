package healthz

import (
	"net/http"

	"strconv"

	"github.com/pkg/errors"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/config"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/database"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/logger"
)

// Manager ...
type Manager struct {
	Client Client
}

var httpNewRequest = http.NewRequest

func (m Manager) getDatabaseStatus(db database.DB) Database {
	result := Database{}
	//Ping database
	err := db.Ping()
	if err != nil {
		result.Error = Error{
			Type:        "Ping",
			Description: "Database health check",
			Error:       errors.Wrap(err, "Healthz database connection").Error(),
		}
	} else {
		result.Status = "OK"
	}
	return result
}

