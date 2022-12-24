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

// main function
func main() {
	config = readConfig()
	scrape(config.Url)

	for _, url := range queuePages {
		scrape(url)
	}

	// log config details
	log.Println(queuePages)
	log.Println(queueProducts)
}
