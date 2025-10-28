package helper

import "google.golang.org/protobuf/encoding/protowire"

func ParseAccounts(data []byte) (name string, email string, photo string, gaia string) {
	pos := 0

	for pos < len(data) {
		fieldNum, wireType, n := protowire.ConsumeTag(data[pos:])

		if n < 0 {
			break
		}

		pos += n

		switch wireType {
		case protowire.BytesType:
			v, b := protowire.ConsumeBytes(data[pos:])

			if b < 0 {
				break
			}

			pos += b

			if fieldNum == 1 {
				return ParseAccounts(v)
			}

			switch fieldNum {
			case 2: // Name
				name = string(v)
			case 3: // Email
				email = string(v)
			case 4: // Photo URL
				photo = string(v)
			case 10: // GAIA ID
				gaia = string(v)
			}
		case protowire.VarintType:
			_, b := protowire.ConsumeVarint(data[pos:])

			if b < 0 {
				break
			}

			pos += b
		default:
			b := protowire.ConsumeFieldValue(fieldNum, wireType, data[pos:])

			if b < 0 {
				break
			}

			pos += b
		}
	}

	return
}
