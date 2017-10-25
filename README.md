The purpose of this repository is to provide an easy way to work with GeoJSON and Protocol Buffer 3 in golang, providing the necessary transformations of the basic geometries and types.

### Compile

```bash
/usr/local/bin/protoc -I=protos protos/geojson.proto --go_out=protos/ protos/*.proto
```