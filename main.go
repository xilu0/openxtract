package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement){
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	} )
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL.String())
	//})
	c.OnHTML("#text-bold", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText(".col-symbol"))
	})
	c.Visit("https://github.com")
	
}
