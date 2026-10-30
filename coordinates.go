package dutchprovinces

import (
	_ "embed"
	"encoding/json"
	"sync"

	"github.com/dylandreimerink/go-rijksdriehoek"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
	"github.com/paulmach/orb/project"
)

var (
	//go:embed rijksdriehoeken.json
	datafile json.RawMessage

	provincesOnce = sync.OnceValue(parseCoordinatesJSON)
)

type province struct {
	Code        string
	Polygon     orb.MultiPolygon
	BoundingBox orb.Bound
}

func parseCoordinatesJSON() []province {
	db, err := geojson.UnmarshalFeatureCollection(datafile)
	if err != nil {
		panic(err)
	}

	var ret []province
	for _, f := range db.Features {
		polygon := project.MultiPolygon(f.Geometry.(orb.MultiPolygon), convertToWSG84)
		prov := province{
			Polygon: polygon,
			// Converting the corners of the RD bounding box isn't enough, because the RD grid is rotated relative to the meridians.
			BoundingBox: polygon.Bound(),
		}
		switch f.Properties.MustString("naam") {
		case "Drenthe":
			prov.Code = "NL-DR"
		case "Flevoland":
			prov.Code = "NL-FL"
		case "Fryslân":
			prov.Code = "NL-FR"
		case "Gelderland":
			prov.Code = "NL-GE"
		case "Groningen":
			prov.Code = "NL-GR"
		case "Limburg":
			prov.Code = "NL-LI"
		case "Noord-Brabant":
			prov.Code = "NL-NB"
		case "Noord-Holland":
			prov.Code = "NL-NH"
		case "Overijssel":
			prov.Code = "NL-OV"
		case "Utrecht":
			prov.Code = "NL-UT"
		case "Zeeland":
			prov.Code = "NL-ZE"
		case "Zuid-Holland":
			prov.Code = "NL-ZH"
		default:
			panic("Unknown province " + f.Properties.MustString("naam"))
		}

		ret = append(ret, prov)
	}

	return ret
}

func convertToWSG84(p orb.Point) orb.Point {
	x, y := rijksdriehoek.RDtoWGS84(p.X(), p.Y())
	return orb.Point{x, y}
}

// LookupLatitudeLongitude returns the ISO 3166-2 province code of the given GPS coordinates.
// The first call needs to prepare the database and will be slow. You can make a call from a startup function to pre-initialize it.
func LookupLatitudeLongitude(lat, lon float64) (string, bool) {
	p := orb.Point{lat, lon}
	var found string
	for _, prov := range provincesOnce() {
		if !prov.BoundingBox.Contains(p) {
			// This check saves 40% off the benchmark.
			continue
		}
		if planar.MultiPolygonContains(prov.Polygon, p) {
			found = prov.Code
			break
		}
	}
	return found, found != ""
}
