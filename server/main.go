package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getHtmlByUrl(url string) string {
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

func getTagsByClass(doc *html.Node, selector string) []*html.Node {
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

	f(doc)

	return result
}

func getTextsByTags(nodes []*html.Node) []string {
	var result []string
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		result = append(result, node.FirstChild.FirstChild.Data)
	}

	return result
}

func main() {
	markup := getHtmlByUrl("http://samurai-tula.ru/bisness-lanchi/")
	doc, _ := html.Parse(strings.NewReader(markup))
	tags := getTagsByClass(doc, "vitrina_header")
	texts := getTextsByTags(tags)
	fmt.Println(texts)
}
