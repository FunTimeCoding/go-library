package system

import (
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestValidateArchitecture(t *testing.T) {
	ValidateArchitecture(constant.AMD64)
}
