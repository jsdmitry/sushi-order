package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const (
	menuItemClass            = "vitrina_element"
	menuItemHeaderClass      = "vitrina_header"
	menuItemImageClass       = "vitrina_image"
	menuItemDescriptionClass = "shopwindow_content"
	menuItemPriceClass       = "wpshop_price"
)

// MenuItem contains caption, image url and description
type MenuItem struct {
	Caption     string
	ImageURL    string
	Description string
	Price       uint64
}

// GetHTMLByURL method return HTML markup by URL
func GetHTMLByURL(url string) string {
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

// GetMenuFromHTML method parse HTML page by URL and return the array of menu items
func GetMenuFromHTML(markup string) []*MenuItem {
	doc, _ := html.Parse(strings.NewReader(markup))
	menuItemsNodes := getNodesBySelector(doc, menuItemClass)

	var result []*MenuItem
	for _, menuItemNode := range menuItemsNodes {
		caption := getCaptionFromNode(menuItemNode)
		imageURL := getImageURLFromNode(menuItemNode)
		description := getDescriptionFromNode(menuItemNode)
		price := getPriceFromNode(menuItemNode)
		menuItem := &MenuItem{caption, imageURL, description, price}
		result = append(result, menuItem)
	}
	return result
}

func getNodeBySelector(parentNode *html.Node, selector string) *html.Node {
	if parentNode != nil {
		result := getNodesBySelector(parentNode, selector)
		if len(result) > 0 {
			return result[0]
		}
	}
	return nil
}

func getNodesBySelector(parentNode *html.Node, selector string) []*html.Node {
	var result []*html.Node
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n != nil {
			if n.Type == html.ElementNode && len(n.Attr) > 0 && n.Attr[0].Key == "class" && n.Attr[0].Val == selector {
				result = append(result, n)
			} else {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
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
	descriptionNode := getNodeBySelector(node, menuItemDescriptionClass)
	if descriptionNode != nil {
		return descriptionNode.FirstChild.Data
	}
	return ""
}

func getPriceFromNode(node *html.Node) uint64 {
	priceNode := getNodeBySelector(node.NextSibling, menuItemPriceClass)
	if priceNode != nil {
		r, _ := regexp.Compile("[0-9]+")
		text := r.FindString(priceNode.FirstChild.Data)
		result, err := strconv.ParseUint(text, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		return result
	}
	return 0
}
