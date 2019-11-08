package rs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/upperwal/go-iudx/util"
)

var (
	baseURL       = "https://pudx.resourceserver.iudx.org.in/resource-server/"
	rsName        = "pscdcl"
	latestVersion = "/v1"
)

type ResourceServerClient struct {
	endpointURL string
}

func NewResourceServerClient() *ResourceServerClient {
	return &ResourceServerClient{
		endpointURL: baseURL + rsName + latestVersion + "/",
	}
}

func (rs *ResourceServerClient) Search(ids []string, q ...util.Query) chan string {
	resChannel := make(chan string, 10)
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			opts := map[string]string{
				"id": id,
			}
			for _, qry := range q {
				for _, p := range qry.Params() {
					keyVal := strings.Split(p, ":")
					opts[keyVal[0]] = keyVal[1]
				}
			}
			data, err := rs.post("search", opts)
			if err != nil {
				wg.Done()
				return
			}
			resChannel <- string(data)
			wg.Done()
		}(id)
	}

	go func() {
		defer close(resChannel)
		wg.Wait()
	}()

	return resChannel
}

func (rs *ResourceServerClient) post(action string, bodyMap interface{}) ([]byte, error) {
	body, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}
	fmt.Println("URL: " + rs.endpointURL + action + "\nBody: " + string(body))
	res, err := http.Post(rs.endpointURL+action, "application/json", bytes.NewBuffer(body))
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
