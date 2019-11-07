package cat

import (
	"github.com/upperwal/go-iudx/rs"
)

type CatalogueItems struct {
	ID string `json:"id"`
}

type Catalogue []CatalogueItems

func (c Catalogue) Latest() chan string {
	ids := make([]string, len(c))
	for _, ct := range c {
		ids = append(ids, ct.ID)
	}

	rsc := rs.NewResourceServerClient()
	return rsc.Search(ids, rs.NewQueryLatest())
}
