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

// InsertMenu method create 'menu' table if not exist and insert menu items to table
func (provider *MySQLDataProvider) InsertMenu(menu []*MenuItem) {
	db := createDBConnection(provider.ConnectionString)
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer tx.Rollback()
	defer db.Close()

	createMenuTable(tx)
	for _, menuItem := range menu {
		insertMenuItem(tx, menuItem)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

func createMenuTable(tx *sql.Tx) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` ("+
		"`id` INT(11) NOT NULL AUTO_INCREMENT, "+
		"`caption` VARCHAR(50) NOT NULL, "+
		"`image_url` VARCHAR(50) NOT NULL, "+
		"`description` VARCHAR(50) NOT NULL, "+
		"`price` INT(11) NOT NULL,PRIMARY KEY (`id`))", menuTableName)

	_, err := tx.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}

func createDBConnection(connectionString string) *sql.DB {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func insertMenuItem(tx *sql.Tx, menuItem *MenuItem) {
	valuesString := fmt.Sprintf(`"%s", "%s", "%s", %d`,
		menuItem.Caption,
		menuItem.ImageURL,
		menuItem.Description,
		menuItem.Price)

	query := fmt.Sprintf(`INSERT INTO %s (caption, image_url, description, price) VALUES(%s)`, menuTableName, valuesString)
	fmt.Println(query)
	_, err := tx.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}
