package main

import (
	"github.com/jsdmitry/sushi-order/server/model"
	"github.com/jsdmitry/sushi-order/server/sql"
)

const (
	connectionString = "root:12qwesdf@/sushi_order"
	commonURL        = "http://samurai-tula.ru"
)

var categoriesRequiredList = []string{"Супы", "Горячие блюда", "Бизнес-ланч"}

func main() {
	categoriesHTML := GetHTMLByURL(commonURL)
	categories := GetCategoriesFromHTML(categoriesHTML, categoriesRequiredList)
	dataProvider := sql.MySQLDataProvider{ConnectionString: connectionString}
	dataProvider.InsertCategories(categories)
	dataProvider.InsertMenuFromCategories(categories, func(url string) []*model.MenuItem {
		menuHTML := GetHTMLByURL(commonURL + url)
		return GetMenuFromHTML(menuHTML)
	})
}
