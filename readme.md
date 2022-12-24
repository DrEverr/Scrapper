# Install
1. clone repo
2. cd to project
3. `go build .`
4. run exe file

# config.json
It's important to first create `config.json` file with Name of shop to scan, Output path for csv to generate, homepage url, regex for category, regex for product, and jquery to extract details about products.

There is example_config.json file for easier start

If we just give query it will extract Text of page, but if we want to extract Attribute Value we need to say whitch attribute we want, by adding `/` after querty to element we want, and we need to provide name of attribute we want, eg. `#my-element/att-value` will extract value of `att-value` in element with `id=my-element`