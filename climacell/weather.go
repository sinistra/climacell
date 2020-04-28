package climacell

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

type FloatValue struct {
	Value *float64 `json:"value"`
	Units string   `json:"units"`
}

type NonNullableTimeValue struct {
	Value time.Time `json:"value"`
}
type Weather struct {
	Lat             float64
	Lon             float64
	Temp            *FloatValue
	ObservationTime NonNullableTimeValue `json:"observation_time"`
}

type LatLon struct{ Lat, Lon float64 }

type ForecastArgs struct {
	// If present, latitude and longitude coordinates we are requesting
	// forecast data for.
	LatLon *LatLon
	// If non-blank, ID for location we are requesting forecast data for.
	LocationID string
	// Unit system to return weather data in. Valid values are "si" and "us",
	// default is "si"
	UnitSystem string
	// Weather data fields we want returned in the response
	Fields []string
	// If nonzero, StartTime indicates the initial timestamp to request weather
	// data from.
	StartTime time.Time
	// If nonzero, EndTime indicates the ending timestamp to request weather
	// data to.
	EndTime time.Time
}

func (args ForecastArgs) QueryParams() url.Values {
	q := make(url.Values)

	if args.LatLon != nil {
		q.Add("lat", strconv.FormatFloat(args.LatLon.Lat, 'f', -1, 64))
		q.Add("lon", strconv.FormatFloat(args.LatLon.Lon, 'f', -1, 64))
	}

	if args.LocationID != "" {
		q.Add("location_id", args.LocationID)
	}
	if args.UnitSystem != "" {
		q.Add("unit_system", args.UnitSystem)
	}

	if len(args.Fields) > 0 {
		q.Add("fields", strings.Join(args.Fields, ","))
	}

	if !args.StartTime.IsZero() {
		q.Add("start_time", args.StartTime.Format(time.RFC3339))
	}
	if !args.EndTime.IsZero() {
		q.Add("end_time", args.EndTime.Format(time.RFC3339))
	}

	return q
}
