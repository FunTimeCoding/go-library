package parse

import (
	"fmt"
	"github.com/prometheus/common/model"
	"log"
)

func VectorFloatSingle(v model.Value) float64 {
	vector, okay := v.(model.Vector)

	if !okay {
		log.Panicf("not a vector")
	}

	count := len(vector)

	if count == 0 {
		return 0
	}

	if count == 1 {
		return float64(vector[0].Value)
	}

	for _, element := range vector {
		fmt.Printf("Sample: %+v\n", element)
	}

	log.Panicf("too many samples: %d", count)

	return -2
}
