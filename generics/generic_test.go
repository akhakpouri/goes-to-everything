package generics

import (
	"testing"
)

var ints = map[string]int64{
	"fist":   5,
	"second": 10,
	"thrid":  15,
	"forth":  20,
	"fifth":  30,
}

var floats = map[string]float64{
	"first":  1.1,
	"second": 2.2,
	"third":  3.3,
	"forth":  44.44,
	"fifth":  55.5,
}

func TestSumInt(t *testing.T) {
	t.Logf("int sums %v", SumInts(ints))
}

func TestSumFloats(t *testing.T) {
	t.Logf("floats sum %v", SumFloats(floats))
}

func TestSumIntsOrFloats(t *testing.T) {
	t.Logf("ints or floats sum %v", SumIntsOrFloats(ints))
	t.Logf("ints or floats sum %v", SumIntsOrFloats(floats))
}
