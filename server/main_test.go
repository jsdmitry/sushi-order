package main

import (
	"testing"

	"github.com/franela/goblin"
)

func TestParseMenu(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Parse menu from markup", func() {
		g.It("Parse menu items", func() {
			markup := `<table><tr><td><div class="vitrina_element"></td><td><div class="vitrina_element"></td></tr><tr><td><div class="vitrina_element"></td></tr></table>`
			menu := GetMenuFromHTML(markup)

			g.Assert(len(menu)).Equal(3)
		})

		g.It("Parse caption, imageURL and description from menu item", func() {
			markup := `<div class="vitrina_element"><div class="vitrina_image"><div><a title="test" href="http://test-image.jpg"></a></div></div><div class="vitrina_header"><a>Test Header</a></div><div class="shopwindow_content">(fish, chicken, rice)</div></div>`
			menu := GetMenuFromHTML(markup)
			menuItem := menu[0]

			g.Assert(menuItem.Caption).Equal("Test Header")
			g.Assert(menuItem.ImageURL).Equal("http://test-image.jpg")
			g.Assert(menuItem.Description).Equal("(fish, chicken, rice)")
		})

		g.It("Parse price by menu item", func() {
			markup := `<table><tr><td><div class="vitrina_element"></div><div><table><tr><td class="wpshop_price">\n\t\t\t\t\t\t320 $.\t\t\t\t\t</td></tr></table></div></td></tr></table>`
			menu := GetMenuFromHTML(markup)
			menuItem := menu[0]

			g.Assert(menuItem.Price).Equal(uint64(320))
		})

		g.It("Parse cation, imageURL and menuURL from category item", func() {
			markup := `<div class="tile"><a href="/sushi/"><span class="title">sushi</span><img alt="sushi" src="http://sushi.png"></a></div>`
			categories := GetCategoriesFromHTML(markup)
			categoryItem := categories[0]

			g.Assert(categoryItem.Caption).Equal("sushi")
			g.Assert(categoryItem.ImageURL).Equal("http://sushi.png")
			g.Assert(categoryItem.MenuURL).Equal("/sushi/")
		})

		g.It("Parse category items from markup", func() {
			markup := `<div><div class="tile"></div><div class="tile"></div></div>`
			categories := GetCategoriesFromHTML(markup)

			g.Assert(len(categories)).Equal(2)
		})
	})
}
