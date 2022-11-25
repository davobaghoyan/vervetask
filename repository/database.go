package repository

import (
	"database/sql"
	"fmt"
)

const driver = "mysql"
const dataSourceName = "@tcp(127.0.0.1:3306)/mydb?parseTime=true"
const userName = "***"
const password = "***"

func GetDb() *sql.DB {
	connectionString := userName + ":" + password + dataSourceName
	db, err := sql.Open(driver, connectionString)

	if err != nil {
		fmt.Println("Connection error", err)
	}

	return db
}
