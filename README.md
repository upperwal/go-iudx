# Go client for India Urban Data Exchange [IUDX]

IUDX is a data exchange platform for Indian smart cities. You can search for available datasets in Pune using the [catalogue UI](https://pudx.catalogue.iudx.org.in).

## Getting started
```sh
go get github.com/upperwal/go-iudx/...
```

## Usage
```go
package main

import (
  "fmt"

  "github.com/upperwal/go-iudx/cat"
)

func main() {
  // Create a catalogue client
  cc, err := cat.NewCatalogueClient()
  if err != nil {
    panic(err)
  }

  /* Catalogue has many query and filter parameters */

  // You can look for specific attributes (name-value pair)
  // Eg: you can use "tags" attribute with values like "aqm", "air quality", "buses" etc 
  // This will only look for items with these tags value
  attrQ := cat.NewQueryAttribute()
  attrQ.Append("tags", "aqm")

  // This is a filter on the output.
  // Filtering with "id" will return all items with only "id" object. 
  filterQ := cat.NewQueryFilter()
  filterQ.Append("id")

  // You can also perform spatial search
  // This query will look for all items within a 1000m radius with the given center.
  circleQ := cat.NewQueryCircleSearch(18.539107, 73.903987, 1000)

  // Finally we perform the search operation by passing all the query options.
  // It would return all items matching all the query parameters.
	items, err := cc.Search(attrQ, filterQ, circleQ)
	if err != nil {
		panic(err)
	}

  // Above "Search" operation will only yield the metadata.
  // "Latest" retrives the latest data of all items returned by the search.
  resCh := items.Latest()

  for res := range resCh {
    fmt.Println(res)
  }
}
```