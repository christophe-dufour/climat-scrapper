package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://www.infoclimat.fr/observations-meteo/archives/11/aout/2017/toulouse-blagnac/07630.html")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".tableau-recap-day tr td").Each(func(i int, s *goquery.Selection) {
		//	fmt.Println(s.Html())
		fmt.Printf("td n%d: %s \n", i, s.Find("div").Text())

		s.Each(func(i int, s *goquery.Selection) {

			fmt.Println(len(s.Nodes))
		})
	})
}

func main() {
	ExampleScrape()
}
