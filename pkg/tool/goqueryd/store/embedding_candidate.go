package store

type embeddingCandidate struct {
	filePath   string
	collection string
	path       string
	title      string
	hash       string
	body       string
	position   int
	vector     []float32
	distance   float64
}
