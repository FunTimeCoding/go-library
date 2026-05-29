package service

import "k8s.io/apimachinery/pkg/types"

type PatchQuery struct {
	ResourceType string
	Name         string
	Namespace    string
	Patch        string
	Type         types.PatchType
	DryRun       bool
}
