package main

import (
	"testing"

	"github.com/franela/goblin"
)

func TestParseMenu(t *testing.T) {
	markup := `
		<!DOCTYPE html>
		<html>
			<body>
				<div>
					<div class="vitrina_element">
						<div class="vitrina_image"><div><a title="test" href="http://test-image.jpg"></a></div></div>
						<div class="vitrina_header"><a>Test Header</a></div>
						<div class="shopwindow_content">(fish, chicken, rice)</div>
					</div>
				</div>
			</body>
		</html>
	`
	g := goblin.Goblin(t)
	g.Describe("Parse menu from markup", func() {
		g.It("Parse caption, imageURL and description", func() {
			menu := GetMenuFromHTML(markup)
			menuItem := menu[0]
			g.Assert(len(menu)).Equal(1)
			g.Assert(menuItem.Caption).Equal("Test Header")
			g.Assert(menuItem.ImageURL).Equal("http://test-image.jpg")
			g.Assert(menuItem.Description).Equal("(fish, chicken, rice)")
		})
	})
}