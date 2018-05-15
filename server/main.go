package main

const (
	connectionString = "root:12qwesdf@/sushi_order"
	bisnessLanchiURL = "http://samurai-tula.ru/bisness-lanchi/"
)

func main() {
	html := GetHTMLByURL(bisnessLanchiURL)
	menu := GetMenuFromHTML(html)
	dataProvider := MySQLDataProvider{ConnectionString: connectionString}
	dataProvider.InsertMenu(menu)
}
