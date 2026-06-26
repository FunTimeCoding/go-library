package zone

type Zone struct {
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Internal         bool     `json:"internal"`
	Disabled         bool     `json:"disabled"`
	DnssecStatus     string   `json:"dnssecStatus"`
	SoaSerial        int      `json:"soaSerial"`
	LastModified     string   `json:"lastModified"`
	Catalog          string   `json:"catalog"`
	NotifyFailed     bool     `json:"notifyFailed"`
	NotifyFailedFor  []string `json:"notifyFailedFor"`
	Expiry           string   `json:"expiry"`
	IsExpired        bool     `json:"isExpired"`
	SyncFailed       bool     `json:"syncFailed"`
	ValidationFailed bool     `json:"validationFailed"`
}
