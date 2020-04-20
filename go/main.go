package main

func main() {
	dbName := "Arilla"
	dbSurName := "Gorilla"
	dbStaffKindID := 6
	dbStaff := 3

	openDB()
	readDB()
	// updateDB(dbName, dbSurName, dbStaffKindID, dbStaff)
	insertDB(dbName, dbSurName, dbStaffKindID, dbStaff)
	readDB()

}
