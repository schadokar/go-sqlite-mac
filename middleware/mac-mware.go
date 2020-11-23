package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3" // import sqlite package
)

type macAddressRecord struct {
	ID      string `json:"mac_id"`
	Address string `json:"mac_address"`
	Group   string `json:"mac_group"`
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%v", err)
	}
}

// CreateMacAddressRecord in the sqlite db
func CreateMacAddressRecord(w http.ResponseWriter, req *http.Request) {
	// get the db connection
	db := getDBConn()

	// close db connection
	defer db.Close()

	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	var params macAddressRecord

	// decode the request parameters to macAddressRecord type
	err := json.NewDecoder(req.Body).Decode(&params)

	checkErr(err)

	// form insert query
	insertQuery := fmt.Sprintf("INSERT INTO %s (mac_address, mac_group) VALUES (?, ?)", dbName)

	// insert into the mac address db
	stmt, err := db.Prepare(insertQuery)

	// close the stmt request
	defer stmt.Close()

	checkErr(err)

	fmt.Println(params.Address, params.Group)
	// execute the insert statement
	res, err := stmt.Exec(params.Address, params.Group)

	checkErr(err)

	// get the insert id
	id, err := res.LastInsertId()

	checkErr(err)

	// return the order created msg in json format
	json.NewEncoder(w).Encode(fmt.Sprintf("Mac Address Record Created with ID: %v", id))
}

// GetAllMacAddressRecords get all the items from the sqlite db
func GetAllMacAddressRecords(w http.ResponseWriter, req *http.Request) {
	// create the db connection
	db := getDBConn()

	// close the db connection
	defer db.Close()

	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	// form get all records query
	selectQuery := fmt.Sprintf("SELECT * FROM %s", dbName)

	// query
	rows, err := db.Query(selectQuery)

	checkErr(err)

	// create an array of macAddressRecord type to store all the macAddressRecords
	var records struct {
		MacAddressRecords []macAddressRecord `json:"records"`
	}

	// iterate over rows to format macAddressRecord in macAddressRecord type
	for rows.Next() {
		var params macAddressRecord

		err = rows.Scan(&params.ID, &params.Address, &params.Group)

		records.MacAddressRecords = append(records.MacAddressRecords, params)
	}

	// return the order array in json format
	json.NewEncoder(w).Encode(records)
}

// GetMacAddressRecord by MacAddress id
func GetMacAddressRecord(w http.ResponseWriter, req *http.Request) {
	// get all the params from the request
	vars := mux.Vars(req)

	// create the db connection
	db := getDBConn()

	// close db connection
	defer db.Close()

	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	// form get all records query
	selectQuery := fmt.Sprintf("SELECT * FROM %s WHERE mac_id=?", dbName)

	// prepare query
	stmtOut, err := db.Prepare(selectQuery)

	// close stmtOut request
	defer stmtOut.Close()

	checkErr(err)

	var record macAddressRecord

	err = stmtOut.QueryRow(vars["macID"]).Scan(&record.ID, &record.Address, &record.Group)

	checkErr(err)

	// return the order in json format
	json.NewEncoder(w).Encode(record)
}

// GetMacAddressesByGroup get all the mac addresses by mac group
func GetMacAddressesByGroup(w http.ResponseWriter, req *http.Request) {
	// get all the params from the request
	vars := mux.Vars(req)

	// create the db connection
	db := getDBConn()

	// close db connection
	defer db.Close()

	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	// form get mac addresses by group query
	selectQuery := fmt.Sprintf("SELECT * FROM %s WHERE mac_group=?", dbName)

	// prepare query
	stmtOut, err := db.Prepare(selectQuery)

	// close stmtOut request
	defer stmtOut.Close()

	checkErr(err)

	// execute get mac addresses by group query
	rows, err := stmtOut.Query(vars["macGroup"])

	checkErr(err)

	// create an array of macAddressRecord type to store all the macAddressRecords
	var records struct {
		MacAddressRecords []macAddressRecord `json:"records"`
	}

	// iterate over rows to format macAddressRecord in macAddressRecord type
	for rows.Next() {

		var record macAddressRecord

		err = rows.Scan(&record.ID, &record.Address, &record.Group)

		records.MacAddressRecords = append(records.MacAddressRecords, record)
	}

	// return the macAddressRecord array in json format
	json.NewEncoder(w).Encode(records)
}

// UpdateMacAddressOfRecord by MacAddress id
func UpdateMacAddressOfRecord(w http.ResponseWriter, req *http.Request) {
	// get all the params from the request
	vars := mux.Vars(req)

	// create the db connection
	db := getDBConn()

	// close db connection
	defer db.Close()

	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	// form update record query
	updateQuery := fmt.Sprintf("UPDATE %s SET mac_address=? WHERE mac_id=?", dbName)

	// prepare query
	stmtOut, err := db.Prepare(updateQuery)

	// close stmtOut request
	defer stmtOut.Close()

	checkErr(err)

	_, err = stmtOut.Exec(vars["macAddress"], vars["macID"])

	checkErr(err)

	// return the order in json format
	json.NewEncoder(w).Encode(fmt.Sprintf("Mac Address of Record Updated for ID: %v", vars["macID"]))
}

// UpdateMacGroupOfRecord by MacAddress id
func UpdateMacGroupOfRecord(w http.ResponseWriter, req *http.Request) {
	// get all the params from the request
	vars := mux.Vars(req)

	// create the db connection
	db := getDBConn()

	// close db connection
	defer db.Close()

	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	// form update record query
	updateQuery := fmt.Sprintf("UPDATE %s SET mac_group=? WHERE mac_id=?", dbName)

	// prepare query
	stmtOut, err := db.Prepare(updateQuery)

	// close stmtOut request
	defer stmtOut.Close()

	checkErr(err)

	_, err = stmtOut.Exec(vars["macGroup"], vars["macID"])

	checkErr(err)

	// return the order in json format
	json.NewEncoder(w).Encode(fmt.Sprintf("Mac Group of Record Updated for ID: %v", vars["macID"]))
}

// DeleteMacAddressRecord by MacAddress id
func DeleteMacAddressRecord(w http.ResponseWriter, req *http.Request) {
	// get all the params from the request
	vars := mux.Vars(req)

	// create the db connection
	db := getDBConn()

	// close db connection
	defer db.Close()

	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	// form delete record query
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE mac_id=?", dbName)

	// prepare query
	stmtOut, err := db.Prepare(deleteQuery)

	// close stmtOut request
	defer stmtOut.Close()

	checkErr(err)

	_, err = stmtOut.Exec(vars["macID"])

	checkErr(err)

	// return the order in json format
	json.NewEncoder(w).Encode(fmt.Sprintf("Mac Address Record deleted with ID: %v", vars["macID"]))
}

// DeleteMacAddressByGroup by Mac Group
func DeleteMacAddressByGroup(w http.ResponseWriter, req *http.Request) {
	// get all the params from the request
	vars := mux.Vars(req)

	// create the db connection
	db := getDBConn()

	// close db connection
	defer db.Close()

	// db name
	dbName := os.Getenv("SQLITE_DB_NAME")

	// form delete record query
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE mac_group=?", dbName)

	// prepare query
	stmtOut, err := db.Prepare(deleteQuery)

	// close stmtOut request
	defer stmtOut.Close()

	checkErr(err)

	res, err := stmtOut.Exec(vars["macGroup"])

	checkErr(err)

	noOfDeletedRecords, err := res.RowsAffected()

	// return the order in json format
	json.NewEncoder(w).Encode(fmt.Sprintf("%v Mac Address Records by Group %s deleted", noOfDeletedRecords, vars["macGroup"]))
}
