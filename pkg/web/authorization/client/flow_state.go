package client

type FlowState struct {
	Verifier    string `json:"verifier"`
	State       string `json:"state"`
	ReturnPath  string `json:"return_path"`
	CallbackLocator string `json:"callback_locator"`
}
