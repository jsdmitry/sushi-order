package main

import (
	"fmt"
)

const bisnessLanchiURL = "http://samurai-tula.ru/bisness-lanchi/"

func main() {
	html := GetHTMLByURL(bisnessLanchiURL)
	menu := GetMenuFromHTML(html)
	fmt.Println(menu)
}
