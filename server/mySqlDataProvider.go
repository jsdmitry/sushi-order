package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	menuTableName     = "menu"
	categoryTableName = "category"
)

// The MySQLDataProvider struct provide CRUD operation for MySQL database
type MySQLDataProvider struct {
	ConnectionString string
}

// InsertMenu method create 'menu' table if not exist and insert menu items to table
func (provider *MySQLDataProvider) InsertMenu(categoryID int, menu []*MenuItem) {
	db := createDBConnection(provider.ConnectionString)
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer tx.Rollback()
	defer db.Close()

	removeAllMenuItems(tx)
	createMenuTable(tx)
	for _, menuItem := range menu {
		insertMenuItem(tx, categoryID, menuItem)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

// InsertCategories method create 'category' table if not exist and insert category items to table
func (provider *MySQLDataProvider) InsertCategories(categories []*CategoryItem) {
	db := createDBConnection(provider.ConnectionString)
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer tx.Rollback()
	defer db.Close()

	createCategoryTable(tx)
	for _, categoryItem := range categories {
		insertCategoryItem(tx, categoryItem)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

func createMenuTable(tx *sql.Tx) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` ("+
		"`id` INT(11) NOT NULL AUTO_INCREMENT, "+
		"`category_id` INT(11) NOT NULL, "+
		"`caption` VARCHAR(50) NOT NULL, "+
		"`image_url` VARCHAR(150) NOT NULL, "+
		"`description` VARCHAR(150) NOT NULL, "+
		"`price` INT(11) NOT NULL,PRIMARY KEY (`id`)) COLLATE='cp1251_general_ci';", menuTableName)

	_, err := tx.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}

func createCategoryTable(tx *sql.Tx) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` ("+
		"`id` INT(11) NOT NULL, "+
		"`caption` VARCHAR(50) NOT NULL, "+
		"`image_url` VARCHAR(150) NOT NULL, "+
		"PRIMARY KEY (`id`)) COLLATE='cp1251_general_ci';", categoryTableName)

	_, err := tx.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}

func removeAllMenuItems(tx *sql.Tx) {
	tx.Exec("DELETE FROM " + menuTableName)
}

func createDBConnection(connectionString string) *sql.DB {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func insertCategoryItem(tx *sql.Tx, categoryItem *CategoryItem) {
	valuesString := fmt.Sprintf(`%d, "%s", "%s"`,
		categoryItem.ID,
		categoryItem.Caption,
		categoryItem.ImageURL)

	query := fmt.Sprintf(`INSERT INTO %s (id, caption, image_url) VALUES(%s)`, categoryTableName, valuesString)
	_, err := tx.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}

func insertMenuItem(tx *sql.Tx, categoryID int, menuItem *MenuItem) {
	valuesString := fmt.Sprintf(`%d, "%s", "%s", "%s", %d`,
		categoryID,
		menuItem.Caption,
		menuItem.ImageURL,
		menuItem.Description,
		menuItem.Price)

	query := fmt.Sprintf(`INSERT INTO %s (category_id, caption, image_url, description, price) VALUES(%s)`, menuTableName, valuesString)
	_, err := tx.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}
