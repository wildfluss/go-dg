package dg

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"os"
)

func ExampleFoo() {
	var fc FeatureCollection

	// 	err := xml.NewDecoder(bytes.NewReader([]byte(`<Foo>
	// 	<featureMembers>
	// 		<Bar>
	// 			<tileMatrix>EPSG:3857:17</tileMatrix>
	// 			<row>50647</row>
	// 			<column>20967</column>
	// 		</Bar>
	// 		<Bar>
	// 			<tileMatrix>EPSG:3857:17</tileMatrix>
	// 			<row>50647</row>
	// 			<column>20967</column>
	// 		</Bar>
	// 	</featureMembers>
	// </Foo>`))).Decode(&fc)
	err := xml.NewDecoder(bytes.NewReader([]byte(str))).Decode(&fc)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", fc)

	// Output:
	// TODO
}

func ExampleGetFeature() {
	client := NewClient(
		nil,
		os.Getenv("CONNECTID"),
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
	)

	tmf, err := client.WebFeatureService.GetFeature(context.Background(), 17, 50647, 20967)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", tmf)

	// Output:
	// TODO

}

const str = `<?xml version="1.0" encoding="UTF-8" ?>
<wfs:FeatureCollection
  xmlns:wfs="http://www.opengis.net/wfs"
  xmlns:xs="http://www.w3.org/2001/XMLSchema"
  xmlns:DigitalGlobe="http://www.digitalglobe.com"
  xmlns:gml="http://www.opengis.net/gml"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  numberOfFeatures="1"
  timeStamp="2020-04-08T18:39:25.236Z"
  xsi:schemaLocation="http://www.opengis.net/wfs https://securewatch.digitalglobe.com/catalogservice/schemas/wfs/1.1.0/wfs.xsd?CONNECTID=4800532d-d825-4887-b35f-91497c3941a1 http://www.digitalglobe.com https://securewatch.digitalglobe.com/catalogservice/wfsaccess?service=WFS&amp;version=1.1.0&amp;request=DescribeFeatureType&amp;typeName=DigitalGlobe%3ATileMatrixFeature&amp;CONNECTID=4800532d-d825-4887-b35f-91497c3941a1"
><gml:featureMembers><DigitalGlobe:TileMatrixFeature
      gml:id="TileFeatureMatrix.EPSG_3857_17_20967_50647"
    ><DigitalGlobe:layer
      >DigitalGlobe:ImageryTileService</DigitalGlobe:layer><DigitalGlobe:tileMatrixSet
      >EPSG:3857</DigitalGlobe:tileMatrixSet><DigitalGlobe:tileMatrix
      >EPSG:3857:17</DigitalGlobe:tileMatrix><DigitalGlobe:row
      >50647</DigitalGlobe:row><DigitalGlobe:column
      >20967</DigitalGlobe:column><DigitalGlobe:tileIdentifier
      >952bef61bd2d20922b339531dbc9cd67</DigitalGlobe:tileIdentifier><DigitalGlobe:featureInTileIdentifier
      >952bef61bd2d20922b339531dbc9cd67:2019-11-01</DigitalGlobe:featureInTileIdentifier><DigitalGlobe:tileWidth
      >256</DigitalGlobe:tileWidth><DigitalGlobe:tileHeight
      >256</DigitalGlobe:tileHeight><DigitalGlobe:geometry><gml:Polygon
          srsDimension="2"
          srsName="urn:x-ogc:def:crs:EPSG:4326"
        ><gml:exterior><gml:LinearRing srsDimension="2"><gml:posList
              >37.80544394 -122.41241453 37.80544394 -122.40966795 37.80761398 -122.40966795 37.80761398 -122.41241453 37.80544394 -122.41241453</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></DigitalGlobe:geometry><DigitalGlobe:features
      ><DigitalGlobe:FinishedFeature
          gml:id="952bef61bd2d20922b339531dbc9cd67"
        ><DigitalGlobe:featureId
          >952bef61bd2d20922b339531dbc9cd67</DigitalGlobe:featureId><DigitalGlobe:geometry
          ><gml:Polygon
              srsDimension="2"
              srsName="urn:x-ogc:def:crs:EPSG:4326"
            ><gml:exterior><gml:LinearRing srsDimension="2"><gml:posList
                  >37.80761398 -122.41241453 37.80544394 -122.41241453 37.80544394 -122.40966795 37.80761398 -122.40966795 37.80761398 -122.41241453</gml:posList></gml:LinearRing></gml:exterior></gml:Polygon></DigitalGlobe:geometry><DigitalGlobe:acquisitionDate
          >2019-11-01 00:00:00</DigitalGlobe:acquisitionDate><DigitalGlobe:sensorType
          >Optical</DigitalGlobe:sensorType><DigitalGlobe:source
          >Multiple</DigitalGlobe:source><DigitalGlobe:sourceUnit
          >Mosaic Product</DigitalGlobe:sourceUnit><DigitalGlobe:productType
          >Natural Color</DigitalGlobe:productType><DigitalGlobe:CE90Accuracy
          >8.4 meters</DigitalGlobe:CE90Accuracy><DigitalGlobe:RMSEAccuracy
          >3.91 meters</DigitalGlobe:RMSEAccuracy><DigitalGlobe:cloudCover
          >0.0</DigitalGlobe:cloudCover><DigitalGlobe:offNadirAngle
          >27.954159</DigitalGlobe:offNadirAngle><DigitalGlobe:sunElevation
          >37.655544000000006</DigitalGlobe:sunElevation><DigitalGlobe:sunAzimuth
          >172.02258</DigitalGlobe:sunAzimuth><DigitalGlobe:groundSampleDistance
          >0.50</DigitalGlobe:groundSampleDistance><DigitalGlobe:groundSampleDistanceUnit
          >Meter</DigitalGlobe:groundSampleDistanceUnit><DigitalGlobe:dataLayer
          >aop</DigitalGlobe:dataLayer><DigitalGlobe:legacyDescription
          /><DigitalGlobe:outputMosaic
          >true</DigitalGlobe:outputMosaic><DigitalGlobe:colorBandOrder
          >RGB</DigitalGlobe:colorBandOrder><DigitalGlobe:assetName
          >FINISHED</DigitalGlobe:assetName><DigitalGlobe:assetType
          >PRODUCT_GEOMETRY</DigitalGlobe:assetType><DigitalGlobe:legacyId
          >VIVID_NA28_19Q4_021230203</DigitalGlobe:legacyId><DigitalGlobe:factoryOrderNumber
          /><DigitalGlobe:perPixelX
          >4.5E-6</DigitalGlobe:perPixelX><DigitalGlobe:perPixelY
          >-4.5E-6</DigitalGlobe:perPixelY><DigitalGlobe:crsFromPixels
          >EPSG:4326</DigitalGlobe:crsFromPixels><DigitalGlobe:url
          /><DigitalGlobe:ageDays
          >159</DigitalGlobe:ageDays><DigitalGlobe:formattedDate
          >2019-11-01</DigitalGlobe:formattedDate><DigitalGlobe:ingestDate
          >2020-01-10 13:31:44</DigitalGlobe:ingestDate><DigitalGlobe:spatialAccuracy
          >1:12,000</DigitalGlobe:spatialAccuracy><DigitalGlobe:earliestAcquisitionDate
          >2019-11-01 00:00:00</DigitalGlobe:earliestAcquisitionDate><DigitalGlobe:latestAcquisitionDate
          >2019-11-01 00:00:00</DigitalGlobe:latestAcquisitionDate><DigitalGlobe:pixelsIngested
          >true</DigitalGlobe:pixelsIngested><DigitalGlobe:preciseGeometry
          >true</DigitalGlobe:preciseGeometry><DigitalGlobe:vendorName
          /><DigitalGlobe:vendorReference /><DigitalGlobe:companyName
          >DigitalGlobe</DigitalGlobe:companyName><DigitalGlobe:acquisitionType
          >Standard</DigitalGlobe:acquisitionType><DigitalGlobe:orbitDirection
          /><DigitalGlobe:licenseType /><DigitalGlobe:isBrowse
          >false</DigitalGlobe:isBrowse><DigitalGlobe:isMirrored
          >false</DigitalGlobe:isMirrored><DigitalGlobe:isMultipleWKB
          >false</DigitalGlobe:isMultipleWKB><DigitalGlobe:copyright
          >Image Copyright 2020 DigitalGlobe Inc</DigitalGlobe:copyright><DigitalGlobe:beamMode
          /><DigitalGlobe:polarisationMode /><DigitalGlobe:polarisationChannel
          /><DigitalGlobe:antennaLookDirection /><DigitalGlobe:niirs
          >0.0</DigitalGlobe:niirs><DigitalGlobe:tagsAsString
          /></DigitalGlobe:FinishedFeature></DigitalGlobe:features></DigitalGlobe:TileMatrixFeature></gml:featureMembers></wfs:FeatureCollection>
`
