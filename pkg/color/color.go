package color

type Color string

const (
	Reset = "\u001b[0m"

	Bold      = "\u001b[1m"
	Underline = "\u001b[4m"
	Inverse   = "\u001b[7m"
	Italic    = "\u001b[3m"

	Black  = "\u001b[30m"
	Blue   = "\u001b[34m"
	Cyan   = "\u001b[36m"
	Green  = "\u001b[32m"
	Purple = "\u001b[35m"
	Red    = "\u001b[31m"
	White  = "\u001b[37m"
	Yellow = "\u001b[33m"

	BrightBlack   = "\u001b[30;1m"
	BrightBlue    = "\u001b[34;1m"
	BrightCyan    = "\u001b[36;1m"
	BrightGreen   = "\u001b[32;1m"
	BrightMagenta = "\u001b[35;1m"
	BrightRed     = "\u001b[31;1m"
	BrightWhite   = "\u001b[37;1m"
	BrightYellow  = "\u001b[33;1m"

	BackgroundBlack   = "\u001b[40m"
	BackgroundBlue    = "\u001b[44m"
	BackgroundCyan    = "\u001b[46m"
	BackgroundGreen   = "\u001b[42m"
	BackgroundMagenta = "\u001b[45m"
	BackgroundRed     = "\u001b[41m"
	BackgroundWhite   = "\u001b[47m"
	BackgroundYellow  = "\u001b[43m"

	BackgroundBrightBlack   = "\u001b[40;1m"
	BackgroundBrightBlue    = "\u001b[44;1m"
	BackgroundBrightCyan    = "\u001b[46;1m"
	BackgroundBrightGreen   = "\u001b[42;1m"
	BackgroundBrightMagenta = "\u001b[45;1m"
	BackgroundBrightRed     = "\u001b[41;1m"
	BackgroundBrightWhite   = "\u001b[47;1m"
	BackgroundBrightYellow  = "\u001b[43;1m"
)
