package record

type Record struct {
	Disabled     bool           `json:"disabled"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	TTL          int            `json:"ttl"`
	Payload      map[string]any `json:"rData"`
	DnssecStatus string         `json:"dnssecStatus"`
	Comments     string         `json:"comments"`
	LastUsedOn   string         `json:"lastUsedOn"`
	LastModified string         `json:"lastModified"`
	ExpiryTTL    int            `json:"expiryTtl"`
}
