package kubernetes

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"log"
)

func ValidateName(s string) {
	if !constant.NameExpression.MatchString(s) {
		log.Panicf("invalid name: %q", s)
	}
}
