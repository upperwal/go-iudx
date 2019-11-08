package cat

import (
	"github.com/upperwal/go-iudx/rs"
)

type CatalogueItems struct {
	ID string `json:"id"`
}

type Catalogue []CatalogueItems

func (c Catalogue) Latest() chan string {
	ids := make([]string, 0, len(c))
	for _, ct := range c {
		if ct.ID == "" {
			continue
		}
		ids = append(ids, ct.ID)
	}

	rsc := rs.NewResourceServerClient()
	return rsc.Search(ids, rs.NewQueryLatest())
}
