package request_context

func (c *Context) ProtocolVersion() string {
	switch c.request.ProtoMajor {
	case 1:
		if c.request.ProtoMinor == 0 {
			return "1.0"
		}

		return "1.1"
	case 2:
		return "2"
	case 3:
		return "3"
	default:
		return ""
	}
}
