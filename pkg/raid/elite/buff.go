package elite

type Buff struct {
	Identifier int         `json:"id"`
	Entries    []BuffEntry `json:"buffData"`
}
