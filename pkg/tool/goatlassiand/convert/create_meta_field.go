package convert

type CreateMetaField struct {
	Name          string              `json:"name"`
	Key           string              `json:"key"`
	Required      bool                `json:"required"`
	Schema        string              `json:"schema"`
	AllowedValues []CreateMetaAllowed `json:"allowed_values,omitempty"`
}
