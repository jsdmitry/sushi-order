package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// The MySQLDataProvider struct provide CRUD operation for MySQL database
type MySQLDataProvider struct {
	ConnectionString string
}

// NewMenuItem method creates a menu item in a database
func (provider *MySQLDataProvider) NewMenuItem(tableName string, menuItem *MenuItem) {
	valuesString := fmt.Sprintf(`%d, "%s", "%s", "%s", %d`,
		1,
		menuItem.Caption,
		menuItem.ImageURL,
		menuItem.Description, 100)

	db := provider.createDBConnection()
	_, err := db.Exec(fmt.Sprintf(`INSERT INTO %s (id, caption, image_url, description, price) VALUES(%s)`, tableName, valuesString))
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
