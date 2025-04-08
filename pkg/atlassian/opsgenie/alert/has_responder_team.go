package alert

import rawAlert "github.com/opsgenie/opsgenie-go-sdk-v2/alert"

func (a *Alert) HasResponderTeam(identifier string) bool {
	for _, r := range a.Responders {
		if r.Type == rawAlert.TeamResponder && r.Id == identifier {
			return true
		}
	}

	return false
}
