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
