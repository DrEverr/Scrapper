package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

// scrape url
func scrape(url string) {
	// get html
	html, err := getHtml(url)
	if err != nil {
		log.Println(err)
	}
	// get all hrefs
	hrefs := getHrefs(html)
	log.Println(hrefs)

	//categories := hrefs.
	// get products
	//products := getProducts(html)
	// get product details
	// for _, product := range products {
	// 	getProductDetails(product)
	// }
}

// main function
func main() {
	config = readConfig()
	scrape(config.Url)
	// log config details
	log.Println(config)
}
