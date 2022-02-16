package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type systemParams struct {
	dbType   string
	host     string
	port     int
	username string
	password string
	dbname   string
}

func connectToPostgres(connect systemParams) (string, *sql.DB, error) {
	// var db sql.DB
	// var err error = nil
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connect.host, connect.port, connect.username, connect.password, connect.dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err.Error(), nil, err
	}

	err = db.Ping()
	if err != nil {
		return err.Error(), nil, err
	}
	return "connected successfully!", db, nil

}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//resp, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	log.Printf("Error reading body: %v", err)
	// 	http.Error(w, "can't read body", http.StatusBadRequest)
	// 	return
	// }
	req.ParseForm()
	port, err := strconv.Atoi(req.FormValue("port"))
	if err != nil {
		log.Printf("Error port was not a number: %v", req.Form)
		http.Error(w, "port not a number", http.StatusBadRequest)
		return
	}

	var connectionInfo = systemParams{
		dbType:   req.FormValue("Database Type"),
		host:     req.FormValue("hname"),
		port:     port,
		username: req.FormValue("uname"),
		password: req.FormValue("pword"),
		dbname:   req.FormValue("dbname"),
	}
	//log.Printf("header = %v\n", req.Header)
	log.Printf("databasetype = %v\n", req.FormValue("Database Type"))
	connectMessage, db, err := connectToPostgres(connectionInfo)
	if err != nil {
		log.Printf("Error port was not a number: %v", req.Form)
		http.Error(w, connectMessage, http.StatusBadRequest)
		return
	}

	var vin string

	err = db.QueryRow("select vin from dbo.car LIMIT 1").Scan(&vin)
	if err != nil {
		log.Printf("Error port was not a number: %v", req.Form)
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	io.WriteString(w, vin)
	defer db.Close()

}
