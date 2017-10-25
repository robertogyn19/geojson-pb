package geojson_pb

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	blankFeature = `{
		"type": "Feature",
		"geometry": null,
		"properties": null
	}`

	invalidFeatureType = `{}`
)

func TestFeature(t *testing.T) {
	fixtures := []struct {
		desc    string
		payload string
		valid   bool
	}{
		{
			desc:    "blank feature",
			payload: blankFeature,
			valid:   true,
		},
		{
			desc:    "invalid feature type",
			payload: invalidFeatureType,
		},
	}

	for _, fix := range fixtures {
		t.Run(fix.desc, func(t *testing.T) {
			feature := new(Feature)
			err := json.Unmarshal([]byte(fix.payload), feature)

			if fix.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
