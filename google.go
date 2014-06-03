// Copyright 2014 Antoine Corcy. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package geocoder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	ENDPOINT = "http://maps.googleapis.com/maps/api/geocode/json?address=%v&sensor=false"
)

// Google defines the Google provider
type Google struct {
	locale string
	region string
	UseSsl bool
}

type Results struct {
	Results []Result `json:"results"`
}

type Result struct {
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
}

type Geometry struct {
	Location Coordinate `json:"location"`
}

func (google *Google) Geocode(a Address) (*Coordinate, error) {
	url := fmt.Sprintf(ENDPOINT, url.QueryEscape(strings.TrimSpace(a.Street)))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var results Results
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	location := results.Results[0].Geometry.Location

	return &Coordinate{location.Lat, location.Lng}, nil
}

func (google *Google) Reverse(c Coordinate) (*Address, error) {
	url := fmt.Sprintf(ENDPOINT, fmt.Sprintf("%f,%f", c.Lat, c.Lng))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var results Results
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return &Address{results.Results[0].FormattedAddress}, nil
}
