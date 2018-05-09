package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	menuTableName = "menu"
)

// The MySQLDataProvider struct provide CRUD operation for MySQL database
type MySQLDataProvider struct {
	ConnectionString string
}

func (provider *MySQLDataProvider) insertMenuItem(menuItem *MenuItem) {
	valuesString := fmt.Sprintf(`%d, "%s", "%s", "%s", %d`,
		1,
		menuItem.Caption,
		menuItem.ImageURL,
		menuItem.Description, 100)

	db := provider.createDBConnection()
	_, err := db.Exec(fmt.Sprintf(`INSERT INTO %s (id, caption, image_url, description, price) VALUES(%s)`, menuTableName, valuesString))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func (provider *MySQLDataProvider) createDBConnection() *sql.DB {
	db, err := sql.Open("mysql", provider.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
