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
  cc, err := cat.NewCatalogueClient()
  if err != nil {
    panic(err)
  }

  attrQ := cat.NewQueryAttribute()
  attrQ.Append("tags", []string{"aqm"})

  filterQ := cat.NewQueryFilter()
  filterQ.Append("id")

  cat, err := cc.Search(attrQ, filterQ)
  if err != nil {
    panic(err)
  }

  resCh := cat.Latest()

  for res := range resCh {
    fmt.Println(res)
  }
}
```