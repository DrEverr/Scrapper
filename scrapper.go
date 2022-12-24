package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var queuePages []string
var queueProducts []string
var extractedProducts []Product
var config Config

// read file config.json and change it to struct
func readConfig() Config {
	var config Config
	// check if file exists
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		log.Fatal("config.json file does not exist")
	}
	// read file
	file, _ := os.ReadFile("config.json")
	json.Unmarshal(file, &config)
	return config
}

// get the contents of a web page with given url
func getHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// get all hrefs from html using jquery
func getHrefs(html string) []string {
	// get all hrefs
	hrefs, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
		return nil
	}

	return hrefs.Find("a").Map(func(_ int, s *goquery.Selection) string {
		href, _ := s.Attr("href")
		return href
	})
}

// testing regex with hrefs using regexp
func compareRegex(hrefs []string, regex string) []string {
	var newHrefs []string
	for _, href := range hrefs {
		matched, err := regexp.MatchString(regex, href)
		if err != nil {
			log.Println(err)
		}
		if matched {
			newHrefs = append(newHrefs, href)
		}
	}
	return newHrefs
}

// scrape url
func scrape(url string) {
	// checking if url contains config.Url
	if !strings.Contains(url, config.Url) {
		url = config.Url + url
	}

	log.Println("Scraping: " + url)

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
	for _, product := range queueProducts {
		log.Println("Extracting: " + product)

		// get html
		html, err := getHtml(product)
		if err != nil {
			log.Println(err)
			continue
		}

		product := extractProductData(html)
		if product.Name != "" && product.Price != "" {
			extractedProducts = append(extractedProducts, product)
		}
	}
}

// extract product data
func extractProductData(html string) Product {
	var product Product

	productHtml, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
		return product
	}

	// get product name
	if config.Product.Name != "" {
		product.Name = find(productHtml, config.Product.Name)
	} else {
		log.Println("Product name not found in config.json")
	}

	// get product price
	if config.Product.Price != "" {
		product.Price = find(productHtml, config.Product.Price)
	} else {
		log.Println("Product price not found in config.json")
	}

	// get product sku
	product.Sku = find(productHtml, config.Product.Sku)

	return product
}

func find(doc *goquery.Document, selector string) string {
	var text string
	if strings.Contains(selector, "/") {
		attribute := strings.Split(selector, "/")[1]
		selector = strings.Split(selector, "/")[0]
		doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
			text, _ = s.Attr(attribute)
		})
	} else {
		doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
	}
	return text
}

// main function
func main() {
	config = readConfig()
	scrape(config.Url)

	for i := 0; i < len(queuePages); i++ {
		scrape(queuePages[i])
	}

	fetchProducts()
}
