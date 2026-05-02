package alert

import "github.com/opsgenie/opsgenie-go-sdk-v2/alert"

func (a *Alert) HasResponderTeam(identifier string) bool {
	for _, r := range a.Responders {
		if r.Type == alert.TeamResponder && r.Id == identifier {
			return true
		}
	}

	return false
}
