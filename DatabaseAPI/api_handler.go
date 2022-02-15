package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func connectToPostgres(connect systemParams) (string, *sql.DB) {
	// var db sql.DB
	// var err error = nil
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connect.host, connect.port, connect.username, connect.password, connect.dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err.Error(), nil
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err.Error(), nil
	}
	return "connected successfully!", db

}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, "Hello, world!\n")
	var connectionInfo = systemParams{
		dbType:   "postgres",
		host:     "localhost",
		port:     5432,
		username: "postgres",
		password: "postgres",
		dbname:   "jauntyj",
	}
	connectMessage, db := connectToPostgres(connectionInfo)
	resp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	log.Printf("formval = %v\n", req.FormValue("Database Type"))
	log.Printf("body = %v\n", string(resp))

	db.Query("select * from dbo.car")

	io.WriteString(w, connectMessage)

}
