package main

const (
	connectionString = "root:12qwesdf@/sushi_order"
	commonURL        = "http://samurai-tula.ru"
)

var categoriesRequiredList = []string{"Супы", "Горячие блюда", "Бизнес-ланч"}

func main() {
	categoriesHTML := GetHTMLByURL(commonURL)
	categories := GetCategoriesFromHTML(categoriesHTML, categoriesRequiredList)
	dataProvider := MySQLDataProvider{ConnectionString: connectionString}
	dataProvider.InsertCategories(categories)
	dataProvider.InsertMenuFromCategories(categories)
}
