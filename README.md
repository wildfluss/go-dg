# go-dg

go-dg is a Go clien library for accessing the [DigitalGlobe API](https://gcs-docs.s3.amazonaws.com/SecureWatch/en-us/Miscellaneous/Misc_DevGuides.htm).

## Usage

```go
import "github.com/ysz/go-dg/dg"
```

Construct a new DigitalGlobe client:

```go
client := NewClient(
    nil,
    os.Getenv("CONNECTID"),
    os.Getenv("USERNAME"),
    os.Getenv("PASSWORD"),
)
```

And access different APIs. For example, get a tile within a bounding box over San Francisco:

```go
tmf, err := client.WebFeatureService.GetFeature(context.Background(), 17, 50647, 20967)

// tmf[0].FeatureInTileIdentifier = 952bef61bd2d20922b339531dbc9cd67:2019-11-01
// tmf[0].TileWidth = 256
// ...
```
