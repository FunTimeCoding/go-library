package store

import "math"

func cosineDistance(
	a []float32,
	b []float32,
) float64 {
	if len(a) != len(b) {
		return 1
	}

	var dot, normA, normB float64

	for i := range a {
		dot += float64(a[i]) * float64(b[i])
		normA += float64(a[i]) * float64(a[i])
		normB += float64(b[i]) * float64(b[i])
	}

	denominator := math.Sqrt(normA) * math.Sqrt(normB)

	if denominator == 0 {
		return 1
	}

	return 1 - dot/denominator
}
