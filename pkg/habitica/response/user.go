package response

type User struct {
	NeedsCron bool      `json:"needsCron"`
	LastCron  string    `json:"lastCron"`
	Stats     Stats     `json:"stats"`
	Items     UserItems `json:"items"`
}
