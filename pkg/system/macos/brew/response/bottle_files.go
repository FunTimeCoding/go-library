package response

type BottleFiles struct {
	Arm64Sequoia  BottleFile `json:"arm64_sequoia,omitempty"`
	Arm64Sonoma   BottleFile `json:"arm64_sonoma,omitempty"`
	Arm64Ventura  BottleFile `json:"arm64_ventura,omitempty"`
	Sonoma        BottleFile `json:"sonoma,omitempty"`
	Ventura       BottleFile `json:"ventura,omitempty"`
	Arm64Linux    BottleFile `json:"arm64_linux,omitempty"`
	X8664Linux    BottleFile `json:"x86_64_linux,omitempty"`
	All           BottleFile `json:"all,omitempty"`
	Arm64Monterey BottleFile `json:"arm64_monterey,omitempty"`
	Arm64BigSur   BottleFile `json:"arm64_big_sur,omitempty"`
	Monterey      BottleFile `json:"monterey,omitempty"`
	BigSur        BottleFile `json:"big_sur,omitempty"`
	Catalina      BottleFile `json:"catalina,omitempty"`
	Arm64Tahoe    BottleFile `json:"arm64_tahoe,omitempty"`
	Sequoia       BottleFile `json:"sequoia,omitempty"`
	Mojave        BottleFile `json:"mojave,omitempty"`
	HighSierra    BottleFile `json:"high_sierra,omitempty"`
	Sierra        BottleFile `json:"sierra,omitempty"`
	ElCapitan     BottleFile `json:"el_capitan,omitempty"`
}
