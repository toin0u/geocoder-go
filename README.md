geocoder-go
===========

Geocoding library in Go.

[![Build Status](https://api.travis-ci.org/toin0u/geocoder-go.svg)](http://travis-ci.org/toin0u/geocoder-go)

Install
-------

    $ go get github.com/toin0u/geocoder-go
    $ go install github.com/toin0u/geocoder-go

Providers
---------

- *Google*: you can define `UseSsl`, `Locale` and `Region`. There are all optinal.
- ...

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

google := geocoder.Google{true, "fr-Fr", "France"}
coodinate := geocoder.Coordinate{48.856614, 2.3522219}

fmt.Println(google.Reverse(coodinate)) // &{94 Quai de l'Hôtel de ville, 75004 Paris, France} <nil>
```

License
-------

geocoder-go is released under the MIT License. See the bundled
[LICENSE](https://github.com/toin0u/geocoder-go/blob/master/LICENSE) file for details.
