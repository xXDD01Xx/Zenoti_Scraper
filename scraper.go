package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	//scraping logic
	c := colly.NewCollector()
	c.Visit("https://en.Wikipedia.org/wiki/Main_Page")

	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error){
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func (r *colly.Response){
		fmt.Println("Page visited: ", r.Request.URL)
	})



	fmt.Println("Hello, World!");
}