package zone

type Zone struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Internal     bool   `json:"internal"`
	DnssecStatus string `json:"dnssecStatus"`
	SoaSerial    int    `json:"soaSerial"`
	Disabled     bool   `json:"disabled"`
}
