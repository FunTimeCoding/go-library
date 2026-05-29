package mute_rule

import "time"

func New(
	reason string,
	message string,
	cluster string,
) *MuteRule {
	return &MuteRule{
		Reason:    reason,
		Message:   message,
		Cluster:   cluster,
		CreatedAt: time.Now(),
	}
}
