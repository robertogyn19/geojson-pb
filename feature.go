package geojson_pb

import (
	"bytes"
	"encoding/json"
	"errors"

	pgeojson "github.com/paulmach/go.geojson"
	"github.com/robertogyn19/geojson-pb/protos"
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
	InvalidFeatureTypeError  = errors.New("Invalid feature type, must be 'Feature'!")
	InvalidGeometryError     = errors.New("Invalid geometry!")
	InvalidGeometryTypeError = errors.New("Invalid geometry type!")
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
		return InvalidFeatureTypeError
	}

	// TODO Allow geometry: null as a valid geometry value

	geometryData, err := json.Marshal(featureAsMap["geometry"])
	if err != nil {
		return err
	}

	geom, err := checkGeometry(geometryData)
	if err != nil {
		return InvalidGeometryError
	}

	err = checkGeometryType(geom)
	if err != nil {
		return err
	}

	return nil
}

func checkGeometry(data []byte) (*pgeojson.Geometry, error) {
	geom := new(pgeojson.Geometry)
	return geom, geom.UnmarshalJSON(data)
}

func checkGeometryType(geom *pgeojson.Geometry) error {
	switch geom.Type {
	case pgeojson.GeometryPoint:
		fallthrough
	case pgeojson.GeometryMultiPoint:
		fallthrough
	case pgeojson.GeometryLineString:
		fallthrough
	case pgeojson.GeometryMultiLineString:
		fallthrough
	case pgeojson.GeometryPolygon:
		fallthrough
	case pgeojson.GeometryMultiPolygon:
		fallthrough
	case pgeojson.GeometryCollection:
		return nil
	}

	return InvalidGeometryTypeError
}

type Geometry geojson.Geometry
