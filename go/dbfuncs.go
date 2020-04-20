package main

import (
	"database/sql"
	"fmt"
	"log"

	. "./tables"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func openDB() {
	db, err := sql.Open("mysql", "test_user:password@tcp(localhost:3306)/duty")
	if err != nil {
		log.Println(err)
	}
	database = db
}

func updateDB(dbName string, dbSurName string, DbStaffKindID int, dbStaff int) {
	update, err := database.Exec("UPDATE db_staff SET db_name = ?, db_surname = ?, db_staffkindId = ? WHERE iddb_staff = ?", dbName, dbSurName, DbStaffKindID, dbStaff)
	if err != nil {
		log.Println(err)
	}
	_ = update
}

func readDB() {
	read, err := database.Query("SELECT * FROM db_staff")
	if err != nil {
		panic(err.Error())
	}

	for read.Next() {
		var line DbStaff
		// line := NewDbStaff()
		err := read.Scan(&line.IddbStaff, &line.DbSurname, &line.DbName, &line.DbStaffKindID)
		// err := read.Scan(&iddb_staff, &db_surname, &db_name, &db_staffKindId)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(line)
	}
}

func insertDB(dbName string, dbSurName string, dbStaffKindID int, dbStaff int) {
	insert, err := database.Query("INSERT INTO db_staff VALUES(?,?,?,?)", dbStaff, dbSurName, dbName, dbStaffKindID)
	if err != nil {
		log.Println(err)
	}
	defer insert.Close()
}
