# Install
1. clone repo
2. `cd scrapper`
3. `go build .`
4. run exe file

# config.json
It's important to first create `config.json` file.
It should look like this:
```json
{
  "name": "Example Shop",
  "url": "https://example.com",
  "output": "output.txt",
  "category_regex": "category_regex_pattern",
  "product_regex": "product_regex_pattern",
  "product": {
    "sku": "123456",
    "name": "Example Product",
    "price": "99.99",
    "old_price": "149.99",
    "currency": "USD",
    "url": "https://example.com/product",
    "availability": "In Stock",
    "description": "This is an example product description.",
    "image": "https://example.com/image.jpg",
    "category": "Example Category",
    "ean": "1234567890123",
    "brand": "Example Brand",
    "stock": "100"
  }
}
```
Best refer to config.go for more info.

** Info **
```
There is example_config.json file for easier start
```

If we just give query it will extract Text of page, but if we want to extract Attribute Value we need to say which attribute we want, by adding `/` after querty to element we want, and we need to provide name of attribute we want, eg. `#my-element/att-value` will extract value of `att-value` in element with `id=my-element`

# ToDo \### Project status
- [ ] Add next page xpath
- [ ] Remove filtering urls
- [ ] Save output to file (csv) more often
- [ ] Extraction attributes
- [ ] Concatenate/create array of values if more elements on same query exists
- [ ] Multithreading
- [ ] Fix bug with https and http in url. Desn's see http as https and vice versa