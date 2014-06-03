geocoder-go
===========

Geocoding library in Go.

Install
-------

    $ go get github.com/toin0u/geocoder-go
    $ go install github.com/toin0u/geocoder-go

Providers
---------

* Google
* ...

Geocode
-------

``` go
import "github.com/toin0u/geocoder-go"

var google geocoder.Google
address := geocoder.Address{"paris"}

fmt.Println(google.Geocode(address)) // &{48.856614 2.3522219} <nil>
```

Reverse
-------

``` go
import "github.com/toin0u/geocoder-go"

var google geocoder.Google
coodinate := geocoder.Coordinate{48.856614, 2.3522219}

fmt.Println(google.Reverse(coodinate)) // &{94 Quai de l'Hôtel de ville, 75004 Paris, France} <nil>
```

License
-------

geocoder-go is released under the MIT License. See the bundled
[LICENSE](https://github.com/toin0u/geocoder-go/blob/master/LICENSE) file for details.
