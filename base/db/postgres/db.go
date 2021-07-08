package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"gopkg.in/pg.v4"
)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func DBConnecting() {
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))
	log.Println("Data:", port)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"), port, os.Getenv("PG_USER"), os.Getenv("PG_PSWD"), os.Getenv("PG_DB"))
	log.Println("PSQL Info:", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	log.Println("DB and Error:", db, err)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func GetClient() *sql.DB {
	return db
}

var PG *pg.DB

func InitPG() *pg.DB {
	PG = pg.Connect(&pg.Options{
		Addr:     os.Getenv("PG_HOST_WITH_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PSWD"),
		Database: os.Getenv("PG_DB"),
	})

	return PG
}
