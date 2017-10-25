package geojson_pb

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	geometryPoint = `{
		"type": "Point",
		"coordinates": [ -46, -23 ]
	}`

	geometryMultiPoint = `{
		"type": "MultiPoint",
		"coordinates": [ [-46.0, -23.1], [ -43.2, -26.3 ] ]
	}`

	geometryPolygonNoHoles = `{
		 "type": "Polygon",
		 "coordinates": [
			 [
				 [100.0, 0.0],
				 [101.0, 0.0],
				 [101.0, 1.0],
				 [100.0, 1.0],
				 [100.0, 0.0]
			 ]
		 ]
	 }`

	geometryPolygonWithHoles = `{
		 "type": "Polygon",
		 "coordinates": [
			 [
				 [100.0, 0.0],
				 [101.0, 0.0],
				 [101.0, 1.0],
				 [100.0, 1.0],
				 [100.0, 0.0]
			 ],
			 [
				 [100.8, 0.8],
				 [100.8, 0.2],
				 [100.2, 0.2],
				 [100.2, 0.8],
				 [100.8, 0.8]
       ]
		 ]
	 }`
)

func TestGeometry(t *testing.T) {
	fixtures := []struct {
		desc    string
		payload string
		valid   bool
	}{
		{
			desc:    "geometry point",
			payload: geometryPoint,
			valid:   true,
		},
		{
			desc:    "geometry multi point",
			payload: geometryMultiPoint,
			valid:   true,
		},
		{
			desc:    "geometry polygon (no holes)",
			payload: geometryPolygonNoHoles,
			valid:   true,
		},
		{
			desc:    "geometry polygon (with holes)",
			payload: geometryPolygonWithHoles,
			valid:   true,
		},
	}

	for _, fix := range fixtures {
		t.Run(fix.desc, func(t *testing.T) {
			geom := new(Geometry)
			err := json.Unmarshal([]byte(fix.payload), geom)

			if fix.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
