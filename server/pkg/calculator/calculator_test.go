package calculator

import (
	"os"
	"testing"

	"github.com/hschendel/stl"
)

func TestCalc(t *testing.T) {
	file, err := os.Open("./Tetrahedron.stl")
	if err != nil {
		panic(err)
	}

	solid, err := stl.ReadAll(file)
	if err != nil {
		panic(err)
	}

	volume := CalculateVolume(solid)
	if int(volume) != 513 {
		t.Errorf("failed my mom %d", int(volume))
	}
}

func TestArea(t *testing.T) {
	file, err := os.Open("./Tetrahedron.stl")
	if err != nil {
		panic(err)
	}

	solid, err := stl.ReadAll(file)
	if err != nil {
		panic(err)
	}

	area := CalculateSurface(solid)
	if int(area) != 461 {
		t.Errorf("failed my mom %f", area)
	}
}

func TestReport(t *testing.T) {
	file, err := os.Open("./Tetrahedron.stl")
	if err != nil {
		panic(err)
	}

	report, err := Calculate(Details{
		TechnologyFilling: 0.80,
		Quality:           0.80,
		PlasticCost:       1.2,
		TechnologyCost:    1.5,
		ServiceCost:       100,
	}, file)
	if err != nil {
		panic(err)
	}
	t.Errorf("failed my mom %v", report)

}
