package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const (
	bisnessLanchiURL = "http://samurai-tula.ru/bisness-lanchi/"

	menuItemClass       = "vitrina_element"
	menuItemHeaderClass = "vitrina_header"
	menuItemImageClass  = "vitrina_image"
)

func getHTMLByURL(url string) string {
	response, err := http.Get(url)
	if err != nil {
		return err.Error()
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err.Error()
		}
		return string(contents)
	}
	return ""
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

func getNodesFromParentNodes(parentNodes []*html.Node, selector string) []*html.Node {
	var result, nodes []*html.Node

	for _, parentNode := range parentNodes {
		nodes = getNodesBySelector(parentNode, selector)
		var node *html.Node
		if len(nodes) > 0 {
			node = nodes[0]
		}
		result = append(result, node)
	}
	return result
}

func getTextsFromNodes(nodes []*html.Node) []string {
	var result []string
	for _, node := range nodes {
		result = append(result, node.FirstChild.FirstChild.Data)
	}

	return result
}

func getAttrValueByKey(attrs []html.Attribute, key string) string {
	for _, attr := range attrs {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getImageURLsFromNodes(nodes []*html.Node) []string {
	var result []string
	for _, node := range nodes {
		imageNode := node.FirstChild.FirstChild
		attrValue := getAttrValueByKey(imageNode.Attr, "href")
		result = append(result, attrValue)
	}

	return result
}

func getMenuItemsInfo(url string) ([]string, []string) {
	markup := getHTMLByURL(url)
	doc, _ := html.Parse(strings.NewReader(markup))
	menuItems := getNodesBySelector(doc, menuItemClass)

	menuItemHeaders := getNodesFromParentNodes(menuItems, menuItemHeaderClass)
	texts := getTextsFromNodes(menuItemHeaders)

	menuItemImages := getNodesFromParentNodes(menuItems, menuItemImageClass)
	imageURLs := getImageURLsFromNodes(menuItemImages)

	return texts, imageURLs
}

func main() {
	texts, imageURLs := getMenuItemsInfo(bisnessLanchiURL)
	fmt.Println(texts)
	fmt.Println(imageURLs)
}
