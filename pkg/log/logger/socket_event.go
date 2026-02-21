package logger

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assistant/message"
	"github.com/funtimecoding/go-library/pkg/errors"
	"log/slog"
	"time"
)

// https://opentelemetry.io/docs/specs/semconv/messaging

// Required
const (
	MessagingSystem = "messaging.system"
	OperationName   = "messaging.operation.name"
)

// Recommended
const (
	MessageIdentifier      = "messaging.message.id"
	ConversationIdentifier = "messaging.message.conversation_id"
	BodySize               = "messaging.message.body.size"
	EnvelopeTime           = "messaging.message.envelope_time"
)

// MessageType Optional
const MessageType = "messaging.message.type"

// HomeAssistantOrigin Home Assistant specific
const HomeAssistantOrigin = "messaging.homeassistant.origin"

func (l *Logger) WebSocket(v *message.Message) {
	if v.Event == nil {
		return
	}

	t, e := time.Parse(time.RFC3339Nano, v.Event.Time)
	errors.PanicOnError(e)

	result := []slog.Attr{
		slog.String(MessagingSystem, "homeassistant"),
		slog.String(OperationName, "receive"),
		slog.String(
			MessageIdentifier,
			fmt.Sprintf("%d", v.Identifier),
		),
		slog.String(MessageType, v.Event.Type),
		slog.String(
			ConversationIdentifier,
			v.Event.Context.Identifier,
		),
		slog.Time(EnvelopeTime, t),
		slog.Int(BodySize, len(v.Event.Raw)),
	}

	if v.Event.Origin != "" {
		result = append(
			result,
			slog.String(HomeAssistantOrigin, v.Event.Origin),
		)
	}

	l.structured.LogAttrs(
		l.context,
		slog.LevelInfo,
		"websocket_event_receive",
		result...,
	)
}
