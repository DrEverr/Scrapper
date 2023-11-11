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

	// Then use this function for each attribute
	product.Name = getProductAttributeRequired(productHtml, config.Product.Name, "Product name", true)
	product.Price = getProductAttributeRequired(productHtml, config.Product.Price, "Product price", true)
	product.Sku = getProductAttributeRequired(productHtml, config.Product.Sku, "Product sku", true)
	product.Old_price = getProductAttribute(productHtml, config.Product.Old_price, "Product old price")
	product.Availability = getProductAttribute(productHtml, config.Product.Availability, "Product availability")
	product.Description = getProductAttribute(productHtml, config.Product.Description, "Product description")
	product.Image = getProductAttribute(productHtml, config.Product.Image, "Product image")
	product.Category = getProductAttribute(productHtml, config.Product.Category, "Product category")
	product.Ean = getProductAttribute(productHtml, config.Product.Ean, "Product ean")
	product.Brand = getProductAttribute(productHtml, config.Product.Brand, "Product brand")
	product.Stock = getProductAttribute(productHtml, config.Product.Stock, "Product stock")
	product.Currency = getProductAttribute(productHtml, config.Product.Currency, "Product currency")

	return product
}

func getProductAttribute(productHtml *goquery.Document, configAttribute, attributeName string) string {
	return getProductAttributeRequired(productHtml, configAttribute, attributeName, false)
}

func getProductAttributeRequired(productHtml *goquery.Document, configAttribute, attributeName string, required bool) string {
	if configAttribute != "" {
		return find(productHtml, configAttribute)
	} else if required {
		log.Printf("%s not found in config.json", attributeName)
		return ""
	} else {
		return ""
	}
}
