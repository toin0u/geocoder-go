// Copyright 2014 Antoine Corcy. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package geocoder manages geocoding & reverse geocoding via extern providers
package geocoder

import (
	"net/http"
)

// Address defines a street address to geocode
type Address struct {
	Street string
}

// Coordinate defines a location by its latitude and longitude to reverse geocode
type Coordinate struct {
	Lat, Lng float64
}

// Provider is an interface for geocoding and reverse geocoding
type Provider interface {
	Geocode(a Address) (*Coordinate, error)
	Reverse(c Coordinate) (*Address, error)
}

func request(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
