package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
)

func fetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func getTitle(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "title" {
		if node.FirstChild != nil {
			return node.FirstChild.Data
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		title := getTitle(child)
		if title != "" {
			return title
		}
	}
	return ""
}

func main() {
	url := "https://golang.org"
	fmt.Println("Fetching URL:", url)

	node, err := fetchHTML(url)
	if err != nil {
		fmt.Println("Erro ao buscar HTML:", err)
		return
	}

	title := getTitle(node)
	fmt.Println("Título da página:", title)
}
