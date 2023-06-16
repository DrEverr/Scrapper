package main

import (
	"log"
)

var queuePages []string
var queueProducts []string
var extractedProducts []Product
var config Config

// scrape url
func scrape(url string) {
	fixUrl(&url)

	log.Println("Link extraction: " + url)

	// get html
	html, err := getHtml(url)
	if err != nil {
		log.Println(err)
		return
	}

	// get all hrefs
	hrefs := getHrefs(html)

	// compare hrefs with category regex
	categories := compareRegex(hrefs, config.Category_regex)

	// enqueue unique categories
	for _, category := range categories {
		if !contains(queuePages, category) {
			queuePages = append(queuePages, category)
		}
	}

	// compare hrefs with product regex
	products := compareRegex(hrefs, config.Product_regex)

	// enqueue unique products
	for _, product := range products {
		if !contains(queueProducts, product) {
			queueProducts = append(queueProducts, product)
		}
	}
}

// extract products
func fetchProducts() {
	log.Println("Fetching products")
	for _, productUrl := range queueProducts {
		fixUrl(&productUrl)
		log.Println("Product extraction: " + productUrl)

		// get html
		html, err := getHtml(productUrl)
		if err != nil {
			log.Println(err)
			continue
		}

		product := extractProductData(html)
		if product.Name != "" && product.Price != "" {
			product.Url = productUrl
			extractedProducts = append(extractedProducts, product)
		}
	}
	log.Println("Products fetched")
}

// main function
func main() {
	log.Println("Starting scrapper...")
	config = readConfig()

	log.Println("Scraping: " + config.Name)

	scrape(config.Url)

	for i := 0; i < len(queuePages); i++ {
		scrape(queuePages[i])
	}

	fetchProducts()

	// write to csv
	writeCsv(extractedProducts)
	log.Println("Scrapper finished")
}
