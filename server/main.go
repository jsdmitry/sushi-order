package main

import "fmt"

const (
	connectionString = "root:12qwesdf@/sushi_order"
	bisnessLanchiURL = "http://samurai-tula.ru/bisness-lanchi/"
)

func main() {
	dataProvider := MySQLDataProvider{ConnectionString: connectionString}
	item := MenuItem{"Kegaras", "image", "Ris, volosatiy salat"}
	dataProvider.NewMenuItem("menu", &item)
	html := GetHTMLByURL(bisnessLanchiURL)
	menu := GetMenuFromHTML(html)
	fmt.Println(menu)
}
