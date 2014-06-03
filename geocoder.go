package geocoder

type Address struct {
	Street string
}

type Coordinate struct {
	Lat, Lng float64
}

type Provider interface {
	Geocode(a Address) (*Coordinate, error)
	Reverse(c Coordinate) (*Address, error)
}
