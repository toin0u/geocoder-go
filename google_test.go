// Copyright 2014 Antoine Corcy. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package geocoder

import (
	"testing"
)

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
