package dg

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type FinishedFeature struct {
	XMLName xml.Name `xml:"FinishedFeature"`

	FeatureId string `xml:"featureId"`

	AcquisitionDate string `xml:"acquisitionDate"`
	// "8.4 meters"
	CE90Accuracy   string  `xml:"CE90Accuracy"`
	RMSEAccuracy   string  `xml:"RMSEAccuracy"`
	CloudCover     float64 `xml:"cloudCover"`
	ColorBandOrder string  `xml:"colorBandOrder"`
}

type TileMatrixFeature struct {
	XMLName xml.Name `xml:"TileMatrixFeature"`

	TileMatrix string `xml:"tileMatrix"`
	Row        int    `xml:"row"`
	Column     int    `xml:"column"`

	TileIdentifier          string `xml:"tileIdentifier"`
	FeatureInTileIdentifier string `xml:"featureInTileIdentifier"`

	TileWidth  int `xml:"tileWidth"`
	TileHeight int `xml:"tileHeight"`

	Features []FinishedFeature `xml:"features>FinishedFeature"`
}

// type FeatureMembers struct {

// 	[]TileMatrixFeature
// }

type FeatureCollection struct {
	XMLName xml.Name `xml:"FeatureCollection"`

	FeatureMembers []TileMatrixFeature `xml:"featureMembers>TileMatrixFeature"`
}

type Client struct {
	// Base URL for API requests
	BaseURL *url.URL

	httpClient *http.Client

	WebFeatureService *WebFeatureService

	username  string
	password  string
	connectid string
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}
	x := u.String() + fmt.Sprintf("&connectid=%s", c.connectid)
	// log.Print(x)
	req, err := http.NewRequest(method, x, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	tmp := []byte(fmt.Sprintf("%s:%s", c.username, c.password))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(tmp)))

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(string(bytes))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	// ioutil.WriteFile("dg.xml", body, 0644)
	// log.Printf("%+v\n", string(body))

	err = xml.NewDecoder(bytes.NewReader(body)).Decode(v)
	return resp, err
}

func NewClient(httpClient *http.Client, connectid, username, password string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{}

	u, err := url.Parse("https://securewatch.digitalglobe.com/catalogservice/")
	if err != nil {
		log.Fatal(err)
	}
	c.WebFeatureService = &WebFeatureService{
		c: Client{
			httpClient: httpClient,
			BaseURL:    u,
			connectid:  connectid,
			username:   username,
			password:   password,
		},
	}

	return c
}

type WebFeatureService struct {
	c Client
}

func (s *WebFeatureService) GetFeature(ctx context.Context, zoomlevel, row, column int) ([]TileMatrixFeature, error) {
	// request=GetFeature&typeName=DigitalGlobe:TileMatrixFeature&SERVICE=WFS&VERSION=1.1.0&connectid=&CQL_FILTER=layer='DigitalGlobe:ImageryTileService'and%20tileMatrixSet%20=%20'EPSG:3857'%20and%20tileMatrix%20=%20'EPSG:3857:17'%20and%20row%20=%2050647%20and%20column%20=%2020967
	v := url.Values{}
	v.Set("request", "GetFeature")
	v.Set("typeName", "DigitalGlobe:TileMatrixFeature")
	v.Set("SERVICE", "WFS")
	v.Set("VERSION", "1.1.0")
	v.Set("CQL_FILTER", fmt.Sprintf("layer='DigitalGlobe:ImageryTileService'and tileMatrixSet = 'EPSG:3857' and tileMatrix = 'EPSG:3857:%d' and row = %d and column = %d",
		zoomlevel,
		row,
		column))

	req, err := s.c.newRequest("GET", "wfsaccess?"+v.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var fc FeatureCollection
	_, err = s.c.do(ctx, req, &fc)

	return fc.FeatureMembers, err
}
