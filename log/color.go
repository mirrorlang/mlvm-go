package log

type Colortext string

const (
	// ANSI color escape codes
	Bold        Colortext = "\033[1m"
	Yellow      Colortext = "\033[33m"
	Cyan        Colortext = "\033[36m"
	Gray        Colortext = "\033[90m"
	Red         Colortext = "\033[31m"
	Blue        Colortext = "\033[34m"
	Pink        Colortext = "\033[35m"
	Green       Colortext = "\033[32m"
	LightRed    Colortext = "\033[91m"
	LightGreen  Colortext = "\033[92m"
	LightYellow Colortext = "\033[93m"
	LightBlue   Colortext = "\033[94m"
	LightPink   Colortext = "\033[95m"
	LightCyan   Colortext = "\033[96m"
	White       Colortext = "\033[97m"
	Black       Colortext = "\033[30m"
	EndColor    Colortext = "\033[39m" // "reset everything"

	maxLineWidth = 80
)
