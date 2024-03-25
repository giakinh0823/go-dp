package main

import (
	"fmt"
	"log"
)

type Context struct {
	url     string
	content string
	data    any
}

type Handler func(*Context) error

func CheckingUrlHandler(c *Context) error {
	fmt.Printf("Checking url: %v\n", c.url)
	return nil
}

func FetchContentHandler(c *Context) error {
	fmt.Printf("Fetching content from url: %v\n", c.url)
	c.content = "some content from url"
	return nil
}

func ExtractDataHandler(c *Context) error {
	fmt.Printf("Extracing data from content: %s\n", c.url)
	c.data = map[string]interface{}{"title": "Apple", "price": 10.8}
	return nil
}

func SaveDataHandler(c *Context) error {
	fmt.Printf("Save data from content: %s\n", c.data)
	return nil
}

type HandlerNode struct {
	hdl  Handler
	next *HandlerNode
}

func (node *HandlerNode) Handle(url string) error {
	ctx := Context{url: url}

	if node == nil || node.hdl == nil {
		return nil
	}

	if err := node.hdl(&ctx); err != nil {
		return err
	}

	return node.next.Handle(url)
}

func NewCrawler(handlers ...Handler) HandlerNode {
	node := HandlerNode{}

	if len(handlers) > 0 {
		node.hdl = handlers[0]
	}

	currentNode := &node

	for i := 1; i< len(handlers); i++{
		currentNode.next = &HandlerNode{hdl: handlers[i]}
		currentNode = currentNode.next
	}

	return node
}

type WebCrawler struct {
	handler HandlerNode
}

func (wc WebCrawler) Crawl(url string)  {
	if err := wc.handler.Handle(url) ; err != nil {
		log.Println(err)
	}
}

func main() {
	WebCrawler{
		handler: NewCrawler(
			CheckingUrlHandler,
			FetchContentHandler,
			ExtractDataHandler,
			SaveDataHandler,
		),
	}.Crawl("https://some-rich-content-website/some-page")
}