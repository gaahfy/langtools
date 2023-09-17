package config

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

var db *sql.DB

func GetDb() (*sql.DB) {
    return db
}

func InitDB() {
    dbHost := SqlHost()
    dbPort := SqlPort()
    dbUser := SqlUsername()
    dbPassword := SqlPassword()
    dbName := SqlDatabase()

    // Construisez la chaîne de connexion PostgreSQL
    dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
    var err error
    db, err = sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal(err)
    }

    // Vérifiez la connexion à la base de données
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
}
