package database

import (
	"database/sql"
	"log"
)

var Dbconn *sql.DB

func SetupConnection() {
	var err error
	Dbconn, err = sql.Open("mysql", "root:Piyush@1234@/doctor")
	if err != nil {
		log.Fatal(err)
	}
}
