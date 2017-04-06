package database

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/ricardo-ch/comsab-app/src/api/config"
)

//DbCompta : Database
var DbCompta *sqlx.DB

const connectionString = "server=%s;port=%s;database=%s;user id=%s;password=%s; connection timeout=5"

var sqlConnect = sqlx.Connect

//Connect ...
func Connect() {
	var err error

	DbCompta, err = sqlConnect("mssql", fmt.Sprintf(connectionString, *config.DbHost, *config.DbPort, *config.DbNameCompta, *config.DbUser, *config.DbPassword))
	if err != nil {
		logger.Log.Error(errors.Wrap(err, "Connect to DbCompta database").Error())
	} else {
		DbCompta = DbNew.Unsafe()
	}
}

//DB ...
type DB struct {
	Pinger Pinger
}

// Ping : Ping the db
func (db DB) Ping() error {

	err := db.Pinger.Ping()
	if err != nil {
		err = errors.Wrap(err, "Ping")
		logger.Log.Error(err.Error())
	}
	return err
}

//Close ...
func Close() {
	DbCompta.Close()
}
