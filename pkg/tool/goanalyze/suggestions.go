package goanalyze

type suggestion struct {
	letters []string
	words   []string
}

var suggestions = map[string]suggestion{
	"url":    {letters: []string{"l"}, words: []string{"locator"}},
	"mcp":    {letters: []string{"c"}, words: []string{"model_context"}},
	"dir":    {letters: []string{"d"}, words: []string{"directory"}},
	"dirs":   {words: []string{"directories"}},
	"src":    {letters: []string{"s"}, words: []string{"source"}},
	"dst":    {letters: []string{"d"}, words: []string{"destination"}},
	"tx":     {letters: []string{"t"}},
	"ctx":    {letters: []string{"x"}},
	"param":  {words: []string{"parameter"}},
	"msg":    {letters: []string{"m"}, words: []string{"message"}},
	"req":    {letters: []string{"r"}, words: []string{"request"}},
	"doc":    {letters: []string{"d"}, words: []string{"document"}},
	"pos":    {letters: []string{"p"}, words: []string{"position"}},
	"buf":    {letters: []string{"b"}, words: []string{"buffer"}},
	"ptr":    {letters: []string{"p"}, words: []string{"pointer"}},
	"addr":   {letters: []string{"a"}, words: []string{"address"}},
	"ref":    {letters: []string{"r"}, words: []string{"reference"}},
	"nav":    {letters: []string{"n"}, words: []string{"navigation"}},
	"prev":   {words: []string{"past"}},
	"decl":   {words: []string{"declaration"}},
	"yaml":   {letters: []string{"m"}, words: []string{"markup"}},
	"xml":    {letters: []string{"m"}, words: []string{"markup"}},
	"html":   {letters: []string{"m"}, words: []string{"markup"}},
	"json":   {letters: []string{"j"}, words: []string{"notation"}},
	"config": {letters: []string{"c"}, words: []string{"configuration"}},
	"cfg":    {letters: []string{"c"}, words: []string{"configuration"}},
	"llm":    {letters: []string{"m"}, words: []string{"model"}},
	"tmp":    {letters: []string{"t"}},
	"href":   {words: []string{"link", "reference", "locator"}},
	"def":    {words: []string{"definition"}},
	"concat": {words: []string{"concatenate"}},
	"obj":    {letters: []string{"o"}, words: []string{"object"}},
	"stmt":   {letters: []string{"s"}, words: []string{"statement"}},
	"var":    {letters: []string{"v"}, words: []string{"variable"}},
	"const":  {letters: []string{"c"}, words: []string{"constant"}},
	"func":   {letters: []string{"f"}, words: []string{"function"}},
	"sig":    {letters: []string{"s"}, words: []string{"signature"}},
	"pkg":    {letters: []string{"p"}, words: []string{"package"}},
	"err":    {letters: []string{"e"}, words: []string{"error"}},
	"val":    {letters: []string{"v"}, words: []string{"value"}},
	"args":   {words: []string{"arguments"}},
	"opt":    {letters: []string{"o"}, words: []string{"option"}},
	"cols":   {words: []string{"columns"}},
	"loc":    {letters: []string{"l"}, words: []string{"location"}},
	"dep":    {letters: []string{"d"}, words: []string{"dependency"}},
}

var noSuggestion = map[string]bool{
	"handler": true,
	"data":    true,
	"info":    true,
}
