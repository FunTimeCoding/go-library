package request

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Boolean bool

func (b *Boolean) UnmarshalJSON(y []byte) error {
	var v bool

	if e := json.Unmarshal(y, &v); e == nil {
		*b = Boolean(v)

		return nil
	}

	var s string

	if e := json.Unmarshal(y, &s); e == nil {
		a, f := strconv.ParseBool(s)

		if f != nil {
			return fmt.Errorf("cannot parse %q as boolean", s)
		}

		*b = Boolean(a)

		return nil
	}

	return fmt.Errorf("cannot unmarshal %s as boolean", string(y))
}
