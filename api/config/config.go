package config

import (
	"flag"
)

var (
	Port                 = flag.String("PORT", "7070", "The micro service port")
	DbHost               = flag.String("DB_HOST", "127.0.0.1", "The databases host name")
	DbPort               = flag.String("DB_PORT", "1433", "The databases host port")
	DbNameDbCompta       = flag.String("DB_NAME_Compta", "Compta", "The DbNew database name")
	DbUser               = flag.String("DB_USER", "nassim", "The database user")
	DbPassword           = flag.String("DB_PASSWORD", "nassim", "The database password")
)

//Get : Gets the service configuration
func Get() {
	flag.Parse()

	// if *Port == "" {
	// 	flag.Set("PORT", "7070")
	// }

	// if *DbPort == "" {
	// 	flag.Set("DB_PORT", "1433")
	// }

	//flag.Set("DB_PASSWORD", os.Getenv("DB_PASSWORD"))

	flag.PrintDefaults()
}
