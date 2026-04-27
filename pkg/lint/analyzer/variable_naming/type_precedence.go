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

	if isNamedType(t, "context", "Context") {
		return precedenceContext
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

	if _, okay := underlying.(*types.Map); okay {
		return precedenceMap
	}

	if _, okay := underlying.(*types.Chan); okay {
		return precedenceChannel
	}

	if s, okay := underlying.(*types.Slice); okay {
		e := s.Elem()

		if p, okay := e.(*types.Pointer); okay {
			e = p.Elem()
		}

		if _, okay := e.Underlying().(*types.Struct); okay {
			return precedenceStructSlice
		}

		if named, okay := e.(*types.Named); okay {
			if _, okay := named.Underlying().(*types.Struct); okay {
				return precedenceStructSlice
			}
		}

		return precedencePrimitiveSlice
	}

	if _, okay := underlying.(*types.Struct); okay {
		return precedenceStruct
	}

	if _, okay := underlying.(*types.Interface); okay {
		return precedenceInterface
	}

	return precedenceUnknown
}
