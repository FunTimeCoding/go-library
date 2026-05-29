package model_context

import "k8s.io/apimachinery/pkg/types"

func resolvePatchType(t string) types.PatchType {
	switch t {
	case "merge":
		return types.MergePatchType
	case "json":
		return types.JSONPatchType
	default:
		return types.StrategicMergePatchType
	}
}
