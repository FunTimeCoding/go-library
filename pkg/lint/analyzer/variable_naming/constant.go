package variable_naming

const (
	precedenceError = iota
	precedenceTestingT
	precedenceTestingB
	precedenceReader
	precedenceWriter
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
