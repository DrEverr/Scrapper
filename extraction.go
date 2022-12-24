package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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
	if config.Product.Sku != "" {
		product.Sku = find(productHtml, config.Product.Sku)
	}

	// get product old price
	if config.Product.Old_price != "" {
		product.Old_price = find(productHtml, config.Product.Old_price)
	}

	// get product availability
	if config.Product.Availability != "" {
		product.Availability = find(productHtml, config.Product.Availability)
	}

	// get product description
	if config.Product.Description != "" {
		product.Description = find(productHtml, config.Product.Description)
	}

	// get product image
	if config.Product.Image != "" {
		product.Image = find(productHtml, config.Product.Image)
	}

	// get product description
	if config.Product.Description != "" {
		product.Description = find(productHtml, config.Product.Description)
	}

	// get product category
	if config.Product.Category != "" {
		product.Category = find(productHtml, config.Product.Category)
	}

	// get product ean
	if config.Product.Ean != "" {
		product.Ean = find(productHtml, config.Product.Ean)
	}

	// get product brand
	if config.Product.Brand != "" {
		product.Brand = find(productHtml, config.Product.Brand)
	}

	// get product stock
	if config.Product.Stock != "" {
		product.Stock = find(productHtml, config.Product.Stock)
	}

	// get product currency
	if config.Product.Currency != "" {
		product.Currency = find(productHtml, config.Product.Currency)
	}

	return product
}
