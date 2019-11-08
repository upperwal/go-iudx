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

	circleQ := cat.NewQueryCircleSearch(18.539107, 73.903987, 1000)

	items, err := cc.Search(attrQ, filterQ, circleQ)
	if err != nil {
		panic(err)
	}

	resCh := items.Latest()

	for res := range resCh {
		fmt.Println(res)
	}
}
