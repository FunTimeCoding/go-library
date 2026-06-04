package palette

const (
	scoreMatch          = 16
	scoreGapStart       = -3
	scoreGapExtension   = -1
	bonusWordBoundary   = 8
	bonusWhitespace     = 10
	bonusCamelCase      = 7
	bonusConsecutiveMin = 4
	bonusFirstCharacter = 2
	minScore            = -1 << 30
)
