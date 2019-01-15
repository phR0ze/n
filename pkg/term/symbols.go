package term

// Symbol to rune table
const (
	// Ascii 0 - 31 plus 127 are control characters
	KeyNUL       = rune(0)
	KeyBackSpace = rune(8)
	KeyTab       = rune(9)
	KeyLineFeed  = rune(10)
	KeyReturn    = rune(13)
	KeyEscape    = rune(27)
	KeyDelete    = rune(127)

	KeySpace       = rune(32)
	KeyExclamation = rune(33)
	KeyQuote       = rune(34)
	KeyHash        = rune(35)
	KeyDollar      = rune(36)
	KeyPercent     = rune(37)
	KeyAmpersand   = rune(38)
	KeyApostrophe  = rune(39)
	KeyLeftParen   = rune(40)
	KeyRightParen  = rune(41)
	KeyAsterisk    = rune(42)
	KeyPlus        = rune(43)
	KeyComma       = rune(44)
	KeyHyphen      = rune(45)
	KeyPeriod      = rune(46)
	KeySlash       = rune(47)
	KeyDigit0      = rune(48)
)
