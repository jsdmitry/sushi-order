package main

const (
	connectionString = "root:12qwesdf@/sushi_order"
	commonURL        = "http://samurai-tula.ru"
)

func main() {
	categoriesHTML := GetHTMLByURL(commonURL)
	categories := GetCategoriesFromHTML(categoriesHTML)
	dataProvider := MySQLDataProvider{ConnectionString: connectionString}
	dataProvider.InsertCategories(categories)
	dataProvider.InsertMenuFromCategories(categories)
}
