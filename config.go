package main

// Config is a struct that contains all the configuration for scrapper
type Config struct {
	// Shop name
	Name string `json:"name"`
	// Shop url
	Url string `json:"url"`
	// Categories regex
	Category_regex string `json:"category_regex"`
	// Products regex
	Product_regex string `json:"product_regex"`
	// Product details
	Product Product `json:"product"`
}

type Product struct {
	// Product name
	Name string `json:"name"`
	// Product sku
	Sku string `json:"sku"`
	// Product price
	Price string `json:"price"`
	// Product old price
	Old_price string `json:"old_price"`
	// Product discount
	Discount string `json:"discount"`
	// Product image
	Image string `json:"image"`
	// Product description
	Description string `json:"description"`
	// Product category
	Category string `json:"category"`
	// Product ean
	Ean string `json:"ean"`
	// Product brand
	Brand string `json:"brand"`
	// Product stock
	Stock string `json:"stock"`
	// Product currency
	Currency string `json:"currency"`
	// Product attributes
	Attributes string `json:"attributes"`
}
