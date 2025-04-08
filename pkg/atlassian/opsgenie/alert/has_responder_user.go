package alert

import rawAlert "github.com/opsgenie/opsgenie-go-sdk-v2/alert"

func (a *Alert) HasResponderUser(name string) bool {
	for _, r := range a.Responders {
		if r.Type == rawAlert.UserResponder && r.Name == name {
			return true
		}
	}

	return false
}
