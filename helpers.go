package appconf

import "math"

const (
	floatMaxUlp  = 2
	floatEpsilon = 1e-9
)

// MergeMaps merges two maps
func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	out := map[K]V{}
	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}

// AlmostEqual conducts a combined epsilon / ULP comparison for float64 numbers
func AlmostEqual(f1 float64, f2 float64) bool {
	if math.IsNaN(f1) || math.IsNaN(f2) || math.IsInf(f1, 0) || math.IsInf(f2, 0) {
		return false
	}
	if (f1 < 0) != (f2 < 0) {
		return false
	}
	b1 := math.Float64bits(f1)
	b2 := math.Float64bits(f2)
	if b1 == b2 {
		return true
	}
	delta := f1 - f2
	if delta <= (floatEpsilon * floatMaxUlp) {
		return true
	}
	if b1 > b2 && (b1-b2 <= floatMaxUlp) {
		return true
	}
	if b2 > b1 && (b2-b1 <= floatMaxUlp) {
		return true
	}
	return false
}
