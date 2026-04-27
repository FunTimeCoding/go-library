package variable_naming

const (
	kindLocal variableKind = iota
	kindParameter
	kindReceiver
)

const (
	precedenceError = iota
	precedenceTestingT
	precedenceTestingB
	precedenceReader
	precedenceWriter
	precedenceContext
	precedenceFile
	precedenceGzipWriter
	precedenceTarWriter
	precedenceString
	precedenceInt
	precedenceFloat
	precedenceBool
	precedenceByte
	precedenceByteSlice
	precedenceMap
	precedenceChannel
	precedenceStructSlice
	precedencePrimitiveSlice
	precedenceStruct
	precedenceInterface
	precedenceUnknown
)
