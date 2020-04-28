package climacell

import (
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
