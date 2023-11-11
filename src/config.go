package main

// Config is a struct that contains all the configuration for scrapper
type Config struct {
  // Shop name
  Name string `json:"name"`
  // Shop url
  Url string `json:"url"`
  // Output file
  Output string `json:"output"`
  // Categories regex
  Category_regex string `json:"category_regex"`
  // Products regex
  Product_regex string `json:"product_regex"`
  // Product details
  Product Product `json:"product"`
}

type Product struct {
  // Product sku
  Sku string `json:"sku"`
  // Product name
  Name string `json:"name"`
  // Product price
  Price string `json:"price"`
  // Product old price
  Old_price string `json:"old_price"`
  // Product currency
  Currency string `json:"currency"`
  // Product url
  Url string
  // Product availability
  Availability string `json:"availability"`
  // Product description
  Description string `json:"description"`
  // Product image
  Image string `json:"image"`
  // Product category
  Category string `json:"category"`
  // Product ean
  Ean string `json:"ean"`
  // Product brand
  Brand string `json:"brand"`
  // Product stock
  Stock string `json:"stock"`
}
