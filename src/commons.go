package main

import (
  "encoding/csv"
  "encoding/json"
  "log"
  "os"
  "regexp"
  "strings"

  "github.com/PuerkitoBio/goquery"
)

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

func find(doc *goquery.Document, selector string) string {
  var text string
  if strings.Contains(selector, "/") {
    attribute := strings.Split(selector, "/")[1]
    selector = strings.Split(selector, "/")[0]
    text = doc.Find(selector).AttrOr(attribute, "")
  } else {
    text = doc.Find(selector).Text()
  }
  return text
}

func fixUrl(url *string) {
  // checking if url contains config.Url
  if !strings.Contains(*url, config.Url) {
    *url = config.Url + *url
  }
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

// check if slice contains string
func contains(slice []string, str string) bool {
  for _, v := range slice {
    if v == str {
      return true
    }
  }
  return false
}

func writeCsv(products []Product) {
  log.Println("Writing csv file")
  // create csv file
  file, err := os.Create(config.Output)
  if err != nil {
    log.Println(err)
    return
  }
  defer file.Close()

  // write csv header
  writer := csv.NewWriter(file)
  defer writer.Flush()

  // write csv header
  err = writer.Write([]string{
    "Sku",
    "Name",
    "Price",
    "Old_price",
    "Currency",
    "Url",
    "Currency",
    "Availability",
    "Description",
    "Image",
    "Category",
    "Ean",
    "Brand",
    "Stock",
  })
  if err != nil {
    log.Println(err)
    return
  }

  // write csv rows
  for _, product := range products {
    err := writer.Write([]string{
      product.Sku,
      product.Name,
      product.Price,
      product.Old_price,
      product.Currency,
      product.Url,
      product.Currency,
      product.Availability,
      product.Description,
      product.Image,
      product.Category,
      product.Ean,
      product.Brand,
      product.Stock,
    })
    if err != nil {
      log.Println(err)
      return
    }
  }

  log.Println("Csv file written")
}
