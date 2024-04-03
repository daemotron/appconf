package appconf

import "math"

const (
	floatMaxUlp = 2
)

// mergeMaps merges two maps
func mergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	out := map[K]V{}
	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}

// contains checks whether a collection contains a certain value
func contains[T comparable](collection []T, value T) bool {
	flag := false
	for _, item := range collection {
		flag = flag || (item == value)
	}
	return flag
}

// getEpsilon computes a matching epsilon for float64
func getEpsilon() float64 {
	var d float64 = 1
	for (1.0 + d/2) != 1.0 {
		d = d / 2
	}
	return d
}

// almostEqual conducts a combined epsilon / ULP comparison for float64 numbers
func almostEqual(f1 float64, f2 float64) bool {
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
	delta := math.Abs(f1 - f2)
	if delta <= (getEpsilon() * floatMaxUlp) {
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
