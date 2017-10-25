package geojson_pb

import (
	pgeojson "github.com/paulmach/go.geojson"
	"github.com/robertogyn19/geojson-pb/protos"
)

type Geometry geojson.Geometry

func (geom *Geometry) UnmarshalJSON(data []byte) error {
	geometry, err := pgeojson.UnmarshalGeometry(data)

	if err != nil {
		return err
	}

	// TODO Add support to all geometry types
	switch geometry.Type {
	case pgeojson.GeometryPoint:
		geom.Type = geojson.Geometry_Point

		coords := new(geojson.Coords)
		coords.Coords = geometry.Point

		geom.Point = coords
	case pgeojson.GeometryMultiPoint:
		geom.Type = geojson.Geometry_Multipoint

		coords := make([]*geojson.Coords, len(geometry.MultiPoint))

		for i, path := range geometry.MultiPoint {
			c1 := new(geojson.Coords)
			c1.Coords = path

			coords[i] = c1
		}

		multiCoords := new(geojson.MultiCoords)
		multiCoords.Coords = coords

		geom.Multipoint = multiCoords
	case pgeojson.GeometryPolygon:
		geom.Type = geojson.Geometry_Polygon

		polygon := make([]*geojson.MultiCoords, len(geometry.Polygon))

		for i, ring := range geometry.Polygon {
			coords := make([]*geojson.Coords, len(ring))

			for j, path := range ring {
				c1 := new(geojson.Coords)
				c1.Coords = path

				coords[j] = c1
			}

			innerRing := new(geojson.MultiCoords)
			innerRing.Coords = coords

			polygon[i] = innerRing
		}

		geom.Polygon = polygon
	default:
		return InvalidGeometryTypeError
	}

	return nil
}
