package response

type StackFrame struct {
	Filename string `json:"filename"`
	AbsPath  string `json:"absPath"`
	Module   string `json:"module"`
	Function string `json:"function"`
	LineNo   int    `json:"lineNo"`
	InApp    bool   `json:"inApp"`
}
