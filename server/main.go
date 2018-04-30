package main

import (
	"fmt"
)

const bisnessLanchiURL = "http://samurai-tula.ru/bisness-lanchi/"

func main() {
	menu := GetMenuByURL(bisnessLanchiURL)
	fmt.Println(menu)
}
