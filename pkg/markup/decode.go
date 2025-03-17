package markup

import "gopkg.in/yaml.v3"

func Decode(
	value string,
	structure any,
) error {
	return yaml.Unmarshal([]byte(value), structure)
}
