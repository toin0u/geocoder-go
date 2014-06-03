// Copyright 2014 Antoine Corcy. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package geocoder

import (
	"testing"
)

func TestEndpoint(t *testing.T) {
	google := Google{true, "", ""}

	endpoint := google.buildEndpoint()
	if endpoint != "https://maps.googleapis.com/maps/api/geocode/json?address=%v&sensor=false" {
		t.Errorf("Wrong endpoint: %v", endpoint)
	}

	google.Locale = "da-DK"
	endpoint = google.buildEndpoint()
	if endpoint != "https://maps.googleapis.com/maps/api/geocode/json?address=%v&sensor=false&language=da-DK" {
		t.Errorf("Wrong endpoint: %v", endpoint)
	}

	google.UseSsl = false
	google.Locale = ""
	google.Region = "Denmark"
	endpoint = google.buildEndpoint()
	if endpoint != "http://maps.googleapis.com/maps/api/geocode/json?address=%v&sensor=false&region=Denmark" {
		t.Errorf("Wrong endpoint: %v", endpoint)
	}
}

func TestGeocode(t *testing.T) {
	var google Google
	address := Address{"Paris, France"}

	coordinate, err := google.Geocode(address)
	if err != nil {
		t.Error(err)
	}

	if coordinate.Lat != 48.856614 || coordinate.Lng != 2.3522219 {
		t.Errorf("Wrong coordinate: %v", coordinate)
	}
}

func TestReverse(t *testing.T) {
	var google Google
	coodinate := Coordinate{48.856614, 2.3522219}

	address, err := google.Reverse(coodinate)
	if err != nil {
		t.Error(err)
	}

	if address.Street != "94 Quai de l'Hôtel de ville, 75004 Paris, France" {
		t.Errorf("Wrong address: %v", address)
	}
}
