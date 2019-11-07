package cat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/upperwal/go-iudx/util"
)

var (
	baseURL       = "https://pudx.catalogue.iudx.org.in"
	basePath      = "/catalogue"
	latestVersion = "/v1"
)

type CatalogueClient struct {
	endpointURL string
}

// NewCatalogue creates a catalogue object
func NewCatalogueClient() (*CatalogueClient, error) {
	return &CatalogueClient{
		endpointURL: baseURL + basePath + latestVersion + "/",
	}, nil
}

func (c *CatalogueClient) Search(q ...util.Query) (Catalogue, error) {
	queryString := ""
	for _, qry := range q {
		queryString += qry.String() + "&"
	}
	data, err := c.get("search", queryString)
	if err != nil {
		return nil, err
	}

	var cat Catalogue
	err = json.Unmarshal(data, &cat)
	if err != nil {
		return nil, err
	}

	return cat, nil
}

func (c *CatalogueClient) get(action, query string) ([]byte, error) {
	fmt.Println("Complete: " + c.endpointURL + action + "?" + query)
	res, err := http.Get(c.endpointURL + action + "?" + query)
	if err != nil {
		println("Err: " + err.Error())
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
