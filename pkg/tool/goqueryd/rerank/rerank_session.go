package rerank

import "github.com/amikos-tech/pure-onnx/ort"

type rerankSession struct {
	inputIDs            []int64
	attentionMask       []int64
	inputIDsTensor      *ort.Tensor[int64]
	attentionMaskTensor *ort.Tensor[int64]
	outputTensor        *ort.Tensor[float32]
	session             *ort.AdvancedSession
}
