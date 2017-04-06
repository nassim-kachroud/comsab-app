package healthz

//HealthStatus global status
type HealthStatus struct {
	Status      int         `json:"status"`
	Databases   Databases   `json:"databases"`
	DownloadPDF DownloadPDF `json:"downloadpdf"`
}

//Databases statuses of all databases
type Databases struct {
	Ready   bool     `json:"-"`
	DbNew   Database `json:"dbnew,omitempty"`
	Invoice Database `json:"invoice,omitempty"`
}

//Database status of a database
type Database struct {
	Status string `json:"status,omitempty"`
	Error  Error  `json:"error,omitempty"`
}

//DownloadPDF status of the downloadpdf.aspx solution
type DownloadPDF struct {
	Ready  bool   `json:"-"`
	URL    string `json:"url,omitempty"`
	Status string `json:"status,omitempty"`
	Error  Error  `json:"error,omitempty"`
}

//Error ...
type Error struct {
	Description string `json:"description,omitempty"`
	Error       string `json:"error,omitempty"`
	Type        string `json:"type,omitempty"`
}
