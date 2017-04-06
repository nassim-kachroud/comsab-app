package healthz

import (
	"net/http"

	"github.com/ricardo-ch/after-sales-invoicing/src/api/database"
)

//HealthzManager interface ...
type HealthzManager interface {
	getDatabaseStatus(db database.DB) Database
	getDownloadPDFStatus() DownloadPDF
}

//Client interface
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}
