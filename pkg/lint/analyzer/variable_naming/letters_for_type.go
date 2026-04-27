package variable_naming

import "go/types"

func lettersForType(t types.Type) []string {
	p := typePrecedence(t)

	switch p {
	case precedenceError:
		return []string{"e", "f", "g", "h"}
	case precedenceTestingT:
		return []string{"t"}
	case precedenceTestingB:
		return []string{"b"}
	case precedenceReader:
		return []string{"r"}
	case precedenceWriter:
		return []string{"w"}
	case precedenceContext:
		return []string{"x"}
	case precedenceFile:
		return lettersFromWord("file")
	case precedenceGzipWriter:
		return []string{"z"}
	case precedenceTarWriter:
		return []string{"t"}
	case precedenceString:
		return lettersFromWord("string")
	case precedenceInt:
		return []string{"i"}
	case precedenceFloat:
		return lettersFromWord("float")
	case precedenceBool:
		return []string{"b"}
	case precedenceByte:
		return lettersFromWord("byte")
	case precedenceByteSlice:
		return lettersFromWord("byte")
	case precedenceMap:
		return lettersFromWord("map")
	case precedenceChannel:
		return []string{"c"}
	case precedenceStructSlice:
		return []string{"v"}
	case precedencePrimitiveSlice:
		return lettersForPrimitiveSlice(t)
	case precedenceStruct:
		letters := lettersFromTypeName(t)

		if letters == nil {
			return lettersFromWord("struct")
		}

		return letters
	case precedenceInterface:
		return lettersFromTypeName(t)
	}

	return nil
}
