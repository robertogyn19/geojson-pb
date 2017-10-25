package geojson_pb

import (
	"encoding/json"
	"bytes"

	"github.com/robertogyn19/geojson-pb/protos"
	"errors"
)

const (
	FeatureTypeStr = "Feature"
)

/*
Feature represents the geojson feature entity:

https://tools.ietf.org/html/rfc7946#page-11
 */
type Feature struct {
	geojson.Feature
	Geometry Geometry `json:"geometry"`
}

var (
	InvalidFeatureType = errors.New("Invalid feature type, must be 'Feature'!")
)

func (feature *Feature) UnmarshalJSON(data []byte) error {
	featureAsMap := make(map[string]interface{})

	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()

	err := d.Decode(&featureAsMap)

	if err != nil {
		return err
	}

	// Check type value
	if featureAsMap["type"] != FeatureTypeStr {
		return InvalidFeatureType
	}

	return nil
}

type Geometry geojson.Geometry
