package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const (
	menuItemClass            = "vitrina_element"
	menuItemHeaderClass      = "vitrina_header"
	menuItemImageClass       = "vitrina_image"
	menuItemDedcriptionClass = "shopwindow_content"
)

// MenuItem contains caption, image url and description
type MenuItem struct {
	Caption     string
	ImageURL    string
	Description string
}

// GetMenuByURL method parse HTML page by URL and return the array of menu items
func GetMenuByURL(url string) []MenuItem {
	markup := getHTMLByURL(url)
	doc, _ := html.Parse(strings.NewReader(markup))
	menuItemsNodes := getNodesBySelector(doc, menuItemClass)

	var result []MenuItem
	for _, menuItemNode := range menuItemsNodes {
		caption := getCaptionFromNode(menuItemNode)
		imageURL := getImageURLFromNode(menuItemNode)
		description := getDescriptionFromNode(menuItemNode)
		menuItem := MenuItem{caption, imageURL, description}
		result = append(result, menuItem)
	}
	return result
}

func getNodeBySelector(parentNode *html.Node, selector string) *html.Node {
	var f func(*html.Node) *html.Node

	f = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "div" &&
			len(n.Attr) > 0 && n.Attr[0].Key == "class" && n.Attr[0].Val == selector {
			return n
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				result := f(c)
				if result != nil {
					return result
				}
			}
		}

		return nil
	}

	return f(parentNode)
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
	menuItemHeaderNode := getNodeBySelector(node, menuItemHeaderClass)
	if menuItemHeaderNode != nil {
		return menuItemHeaderNode.FirstChild.FirstChild.Data
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
	imageNode := getNodeBySelector(node, menuItemImageClass)
	if imageNode != nil {
		return getAttrValueByKey(imageNode.FirstChild.FirstChild.Attr, "href")
	}
	return ""
}

func getDescriptionFromNode(node *html.Node) string {
	descriptionNode := getNodeBySelector(node, menuItemDedcriptionClass)
	if descriptionNode != nil {
		return descriptionNode.FirstChild.Data
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
