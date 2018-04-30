package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const (
	menuItemClass       = "vitrina_element"
	menuItemHeaderClass = "vitrina_header"
	menuItemImageClass  = "vitrina_image"
)

// MenuItem contains caption, image url and description
type MenuItem struct {
	Caption     string
	ImageURL    string
	Description string
}

// GetMenuByURL method return array of menu item by URL
func GetMenuByURL(url string) []MenuItem {
	markup := getHTMLByURL(url)
	doc, _ := html.Parse(strings.NewReader(markup))
	menuItemsNodes := getNodesBySelector(doc, menuItemClass)

	var result []MenuItem
	for _, menuItemNode := range menuItemsNodes {
		caption := getCaptionFromNode(menuItemNode)
		imageURL := getImageURLFromNode(menuItemNode)
		menuItem := MenuItem{Caption: caption, ImageURL: imageURL}
		result = append(result, menuItem)
	}
	return result
}

func getNodesBySelector(parentNode *html.Node, selector string) []*html.Node {
	var result []*html.Node
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" &&
			len(n.Attr) > 0 && n.Attr[0].Key == "class" && n.Attr[0].Val == selector {
			result = append(result, n)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}

	f(parentNode)

	return result
}

func getCaptionFromNode(node *html.Node) string {
	menuItemHeaderNodes := getNodesBySelector(node, menuItemHeaderClass)
	if len(menuItemHeaderNodes) > 0 {
		return menuItemHeaderNodes[0].FirstChild.FirstChild.Data
	}
	return ""
}

func getAttrValueByKey(attrs []html.Attribute, key string) string {
	for _, attr := range attrs {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getImageURLFromNode(node *html.Node) string {
	imageNodes := getNodesBySelector(node, menuItemImageClass)
	if len(imageNodes) > 0 {
		return getAttrValueByKey(imageNodes[0].FirstChild.FirstChild.Attr, "href")
	}
	return ""
}

func getHTMLByURL(url string) string {
	var result string
	response, err := http.Get(url)
	if err != nil {
		result = err.Error()
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			result = err.Error()
		}
		result = string(contents)
	}
	return result
}
