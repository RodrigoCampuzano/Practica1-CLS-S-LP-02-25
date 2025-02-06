package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

func ConnectionDB() *sql.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    user := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    host := os.Getenv("HOST")
    database := os.Getenv("DATABASE")
    port := os.Getenv("PORT")

    // a conexión "root:Masterrod123@tcp(localhost:3306)/dbhex"
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error al abrir la base de datos: %v", err)
    }
    if err = db.Ping(); err != nil {
        log.Fatalf("Error al hacer ping a la base de datos: %v", err)
    }

    return db
}