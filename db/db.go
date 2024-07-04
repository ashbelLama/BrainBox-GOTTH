package database

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("mysql", "brainbox:password@tcp(127.0.0.1:6001)/brainbox")
    if err != nil {
        log.Fatalf("error %s when opening DB", err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatal("error connecting to database: ", err)
    }
}
