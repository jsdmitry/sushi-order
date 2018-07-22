package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jsdmitry/sushi-order/server/model"
)

const (
	menuTableName     = "menu"
	categoryTableName = "category"
)

// The SQLDataProvider struct provide CRUD operation for MySQL database
type SQLDataProvider struct {
	ConnectionString string
	db               *sql.DB
}

// ConnectToDB method make connection to the data base
func (provider *SQLDataProvider) ConnectToDB() {
	db, err := sql.Open("mysql", provider.ConnectionString)
	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
	provider.db = db
}

// InsertMenuFromCategories insert a menu items to 'menu' table by categories
func (provider *SQLDataProvider) InsertMenuFromCategories(categories []*model.CategoryItem, getMenuData func(url string) []*model.MenuItem) {
	tx := provider.createTX()
	defer tx.Rollback()

	removeAllItems(tx, menuTableName)
	createMenuTable(tx)

	for _, category := range categories {
		insertMenu(tx, category.ID, getMenuData(category.MenuURL))
	}

	err := tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

// InsertCategories method create 'category' table if not exist and insert category items to table
func (provider *SQLDataProvider) InsertCategories(categories []*model.CategoryItem) {
	tx := provider.createTX()
	defer tx.Rollback()

	removeAllItems(tx, categoryTableName)
	createCategoryTable(tx)
	for _, categoryItem := range categories {
		insertCategoryItem(tx, categoryItem)
	}

	err := tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

// GetCategories method make the get request for the categories data from a server
func (provider *SQLDataProvider) GetCategories() []*model.CategoryItem {
	rows, err := provider.db.Query(fmt.Sprintf("SELECT * FROM %s ORDER BY caption", categoryTableName))

	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	categories := make([]*model.CategoryItem, 0)
	for rows.Next() {
		categoryItem := new(model.CategoryItem)
		err = rows.Scan(&categoryItem.ID, &categoryItem.Caption, &categoryItem.ImageURL)
		if err != nil {
			log.Fatalln(err)
		}
		categories = append(categories, categoryItem)
	}
	return categories
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

func removeAllItems(tx *sql.Tx, tableName string) {
	tx.Exec("DELETE FROM " + tableName)
}

func (provider *SQLDataProvider) createTX() *sql.Tx {
	tx, err := provider.db.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	return tx
}

func insertCategoryItem(tx *sql.Tx, categoryItem *model.CategoryItem) {
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

func insertMenuItem(tx *sql.Tx, categoryID int, menuItem *model.MenuItem) {
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

func insertMenu(tx *sql.Tx, categoryID int, menu []*model.MenuItem) {
	for _, menuItem := range menu {
		insertMenuItem(tx, categoryID, menuItem)
	}
}
