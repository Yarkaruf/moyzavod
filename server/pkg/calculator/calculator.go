package calculator

import (
	"io"
	"math"

	"github.com/golang/geo/r3"
	"github.com/hschendel/stl"
)

// Details of order
type Details struct {
	TechnologyFilling float32
	Quality           float32
	PlasticCost       float32
	TechnologyCost    float32
	ServiceCost       float32
}

// Report ...
type Report struct {
	// Volume of model in mm^3
	Volume float32 `json:"volume"`
	// Surface area of model in cm^2
	Area float32 `json:"area"`
	// In rubles
	Cost   float32 `json:"cost"`
	Bounds Bounds  `json:"bounds"`
}

// Bounds ...
type Bounds struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

// Calculate does report about model
func Calculate(details Details, model io.ReadSeeker) (Report, error) {
	solid, err := stl.ReadAll(model)
	if err != nil {
		return Report{}, err
	}

	volume := CalculateVolume(solid)
	area := CalculateSurface(solid)
	cost := CalculateCost(details, volume)
	l := solid.Measure().Len

	return Report{
		Volume: volume,
		Area:   area,
		Cost:   cost,
		Bounds: Bounds{
			X: l[0],
			Y: l[1],
			Z: l[2],
		},
	}, nil
}

// CalculateCost ...
func CalculateCost(details Details, volume float32) float32 {
	return volume*details.TechnologyFilling*details.TechnologyCost/details.Quality + volume*details.TechnologyFilling*details.PlasticCost + details.ServiceCost
}

// CalculateSurface returns area in cm^2 float32
func CalculateSurface(solid *stl.Solid) float32 {
	var area float64

	for _, poly := range solid.Triangles {
		pa := poly.Vertices[0]
		pb := poly.Vertices[1]
		pc := poly.Vertices[2]

		av := r3.Vector{X: float64(pa[0]), Y: float64(pa[1]), Z: float64(pa[2])}
		bv := r3.Vector{X: float64(pb[0]), Y: float64(pb[1]), Z: float64(pb[2])}
		cv := r3.Vector{X: float64(pc[0]), Y: float64(pc[1]), Z: float64(pc[2])}

		al := av.Distance(bv)
		bl := bv.Distance(cv)
		cl := cv.Distance(av)

		s := (al + bl + cl) / 2
		area += math.Sqrt(s * (s - al) * (s - bl) * (s - cl))
	}

	return float32(area)
}

// CalculateVolume returns volume in mm^3 float32
func CalculateVolume(solid *stl.Solid) float32 {
	var volume float32

	for _, poly := range solid.Triangles {
		pa := poly.Vertices[0]
		pb := poly.Vertices[1]
		pc := poly.Vertices[2]
		v321 := pc[0] * pb[1] * pa[2]
		v231 := pb[0] * pc[1] * pa[2]
		v312 := pc[0] * pa[1] * pb[2]
		v132 := pa[0] * pc[1] * pb[2]
		v213 := pb[0] * pa[1] * pc[2]
		v123 := pa[0] * pb[1] * pc[2]
		volume += (1.0 / 6.0) * (-v321 + v231 + v312 - v132 - v213 + v123)
	}

	return volume
}
