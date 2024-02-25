package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

// example structure
//
//	{
//		3320: {
//			'Country': 'DE',
//			'Name': 'Deutsche Telekom AG',
//			'Routes v4': 13547,
//			'Routes v6': 268
//		},
//		36375: {
//			'Country': 'US',
//			'Name': 'University of Michigan',
//			'Routes v4': 14,
//			'Routes v6': 1
//		}
//	}
type ASNInfo struct {
	Country  string
	Name     string
	RoutesV4 int
	RoutesV6 int
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("bgp.he.net"),
		colly.Async(true),
		colly.MaxDepth(1),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})
	var hrefs []string
	c.OnHTML("a[href*='country']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//push lint to hrefs
		hrefs = append(hrefs, link)
		fmt.Println(link)
		// get all the links and print them
		e.DOM.Find("a[href]").Each(func(_ int, el *goquery.Selection) {
			asnLink, _ := el.Attr("href") // Fix: Use the correct number of variables
			fmt.Println(asnLink)
		})
	})

	start := "https://bgp.he.net/report/world"
	c.Visit(start)

	// visit each link in hrefs and get the data
	// for _, link := range hrefs {
	// 	// c.Visit("https://bgp.he.net" + link)
	// 	// split link to get the country code
	// 	country := link[9:]
	// 	fmt.Println(country)

	// }
	//map the data to the structure

	// save to json file
	// fmt.Println(hrefs)
	c.Wait()
}
