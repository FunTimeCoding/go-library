package forbidden_call

var banned = map[string]bool{
	"Command":        true,
	"CommandContext": true,
}
