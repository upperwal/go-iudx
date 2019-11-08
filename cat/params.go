package cat

import (
	"fmt"
	"strings"
)

type QueryAttribute struct {
	name  []string
	value []string
}

func NewQueryAttribute() *QueryAttribute {
	return &QueryAttribute{}
}

func (qa *QueryAttribute) Append(name string, value ...string) {
	qa.name = append(qa.name, name)
	qa.value = append(qa.value, value...)
}

func (qa *QueryAttribute) Params() []string {
	if len(qa.name) == 0 {
		return nil
	}

	name := "attribute-name=(" + strings.Join(qa.name[:], ",") + ")"
	value := "attribute-value=(" + strings.Join(qa.value[:], ",") + ")"

	return []string{name, value}
}

type QueryFilter struct {
	attributes []string
}

func NewQueryFilter() *QueryFilter {
	return &QueryFilter{}
}

func (qf *QueryFilter) Append(attributes string) {
	qf.attributes = append(qf.attributes, attributes)
}

func (qf *QueryFilter) Params() []string {
	if len(qf.attributes) == 0 {
		return nil
	}

	attr := "attribute-filter=(" + strings.Join(qf.attributes[:], ",") + ")"

	return []string{attr}
}

type QueryCircleSearch struct {
	lat    float32
	lon    float32
	radius uint
}

func NewQueryCircleSearch(lat, lon float32, radius uint) *QueryCircleSearch {
	return &QueryCircleSearch{
		lat:    lat,
		lon:    lon,
		radius: radius,
	}
}

func (qcs *QueryCircleSearch) Params() []string {
	return []string{
		"lat=" + fmt.Sprintf("%f", qcs.lat),
		"lon=" + fmt.Sprintf("%f", qcs.lon),
		"radius=" + fmt.Sprintf("%d", qcs.radius),
	}
}
