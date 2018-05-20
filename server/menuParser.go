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
	menuItemCSSClass            = "vitrina_element"
	menuItemHeaderCSSClass      = "vitrina_header"
	menuItemImageCSSClass       = "vitrina_image"
	menuItemDescriptionCSSClass = "shopwindow_content"
	menuItemPriceCSSClass       = "wpshop_price"
	categoryItemCSSClass        = "tile"
	categoryItemTitleCSSClass   = "title"
)

var categoriesRequiredList = [3]string{"Супы", "Горячие блюда", "Бизнес-ланч"}

// MenuItem contains Caption, ImageUrl and Description
type MenuItem struct {
	Caption     string
	ImageURL    string
	Description string
	Price       uint64
}

// CategoryItem contains Caption, ImageUrl and MenuUrl
type CategoryItem struct {
	ID       int
	Caption  string
	ImageURL string
	MenuURL  string
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

// GetMenuFromHTML method parse HTML page and return the array of menu items
func GetMenuFromHTML(markup string) []*MenuItem {
	doc, _ := html.Parse(strings.NewReader(markup))
	menuItemsNodes := getNodesBySelector(doc, menuItemCSSClass)

	var result []*MenuItem
	for _, menuItemNode := range menuItemsNodes {
		caption := getCaptionFromNode(menuItemNode)
		imageURL := getImageURLFromNode(menuItemNode)
		description := getDescriptionFromNode(menuItemNode)
		price := getPriceFromNode(menuItemNode)
		menuItem := &MenuItem{Caption: caption, ImageURL: imageURL, Description: description, Price: price}
		result = append(result, menuItem)
	}
	return result
}

// GetCategoriesFromHTML method parse HTML page and return array of category items
func GetCategoriesFromHTML(markup string) []*CategoryItem {
	var result []*CategoryItem
	doc, _ := html.Parse(strings.NewReader(markup))
	tileNodes := getNodesBySelector(doc, categoryItemCSSClass)

	for index, tileNode := range tileNodes {
		titleNode := getNodeBySelector(tileNode, categoryItemTitleCSSClass)
		if titleNode != nil {
			caption := titleNode.FirstChild.Data
			if isCategoryCaptionValid(categoriesRequiredList, caption) {
				imageURL := getAttrValueByKey(titleNode.NextSibling.Attr, "src")
				menuURL := getAttrValueByKey(titleNode.Parent.Attr, "href")
				categoryItem := &CategoryItem{index, caption, imageURL, menuURL}
				result = append(result, categoryItem)
			}
		}
	}
	return result
}

func isCategoryCaptionValid(arr [3]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
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
	menuItemHeaderNode := getNodeBySelector(node, menuItemHeaderCSSClass)
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
	imageNode := getNodeBySelector(node, menuItemImageCSSClass)
	if imageNode != nil {
		return getAttrValueByKey(imageNode.FirstChild.FirstChild.Attr, "href")
	}
	return ""
}

func getDescriptionFromNode(node *html.Node) string {
	descriptionNode := getNodeBySelector(node, menuItemDescriptionCSSClass)
	if descriptionNode != nil {
		return strings.Replace(descriptionNode.FirstChild.Data, "½", "1/2", -1)
	}
	return ""
}

func getPriceFromNode(node *html.Node) uint64 {
	priceNode := getNodeBySelector(node.NextSibling, menuItemPriceCSSClass)
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
