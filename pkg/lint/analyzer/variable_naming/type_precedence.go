package variable_naming

import "go/types"

func typePrecedence(t types.Type) int {
	if isErrorType(t) {
		return precedenceError
	}

	if isNamedType(t, "testing", "T") {
		return precedenceTestingT
	}

	if isNamedType(t, "testing", "B") {
		return precedenceTestingB
	}

	if implementsInterface(t, "io", "Reader") {
		return precedenceReader
	}

	if implementsInterface(t, "io", "Writer") {
		return precedenceWriter
	}

	if isNamedType(t, "os", "File") {
		return precedenceFile
	}

	if isNamedType(t, "compress/gzip", "Writer") {
		return precedenceGzipWriter
	}

	if isNamedType(t, "archive/tar", "Writer") {
		return precedenceTarWriter
	}

	underlying := deref(t).Underlying()

	if isBasicKind(underlying, types.String) {
		return precedenceString
	}

	if isIntegerType(underlying) {
		return precedenceInt
	}

	if isFloatType(underlying) {
		return precedenceFloat
	}

	if isBasicKind(underlying, types.Bool) {
		return precedenceBool
	}

	if isByteSlice(t) {
		return precedenceByteSlice
	}

	if isBasicKind(underlying, types.Byte) {
		return precedenceByte
	}

	if _, ok := underlying.(*types.Map); ok {
		return precedenceMap
	}

	if _, ok := underlying.(*types.Chan); ok {
		return precedenceChannel
	}

	if s, ok := underlying.(*types.Slice); ok {
		if _, ok := s.Elem().Underlying().(*types.Struct); ok {
			return precedenceStructSlice
		}

		if named, ok := s.Elem().(*types.Named); ok {
			if _, ok := named.Underlying().(*types.Struct); ok {
				return precedenceStructSlice
			}
		}

		return precedencePrimitiveSlice
	}

	if _, ok := underlying.(*types.Struct); ok {
		return precedenceStruct
	}

	if _, ok := underlying.(*types.Interface); ok {
		return precedenceInterface
	}

	return precedenceUnknown
}
