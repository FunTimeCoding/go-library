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
	case precedenceFile:
		return lettersFromWord("file")
	case precedenceGzipWriter:
		return []string{"z"}
	case precedenceTarWriter:
		return []string{"t"}
	case precedenceString:
		return []string{"s"}
	case precedenceInt:
		return []string{"i"}
	case precedenceFloat:
		return []string{"n"}
	case precedenceBool:
		return []string{"b"}
	case precedenceByte:
		return []string{"b"}
	case precedenceByteSlice:
		return []string{"b"}
	case precedenceMap:
		return []string{"m"}
	case precedenceChannel:
		return []string{"c"}
	case precedenceStructSlice:
		return []string{"v"}
	case precedencePrimitiveSlice:
		return lettersForPrimitiveSlice(t)
	case precedenceStruct:
		return lettersFromTypeName(t)
	case precedenceInterface:
		return lettersFromTypeName(t)
	}

	return nil
}
