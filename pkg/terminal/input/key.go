package input

type Key uint8

const (
	Down Key = iota + 1
	Left
	Right
	Up
	End
	Home
	Insert
	PageDown
	PageUp
)

const Control rune = 0x1f

const (
	// Control Characters
	Null                   rune = 0
	StartOfHeading         rune = 1
	StartOfText            rune = 2
	EndOfText              rune = 3
	EndOfTransmission      rune = 4
	Enquiry                rune = 5
	Acknowledgment         rune = 6
	Bell                   rune = 7
	Backspace              rune = 8
	HorizontalTab          rune = 9
	LineFeed               rune = 10
	VerticalTab            rune = 11
	FormFeed               rune = 12
	CarriageReturn         rune = 13
	ShiftOut               rune = 14
	ShiftIn                rune = 15
	DataLineEscape         rune = 16
	DeviceControl1         rune = 17
	DeviceControl2         rune = 18
	DeviceControl3         rune = 19
	DeviceControl4         rune = 20
	NegativeAcknowledgment rune = 21
	SynchronousIdle        rune = 22
	EndOfTransmitBlock     rune = 23
	Cancel                 rune = 24
	EndOfMedium            rune = 25
	Substitute             rune = 26
	Escape                 rune = 27
	FileSeparator          rune = 28
	GroupSeparator         rune = 29
	RecordSeparator        rune = 30
	UnitSeparator          rune = 31

	// Printable Characters
	Space            rune = 32
	ExclamationMark  rune = 33
	DoubleQuote      rune = 34
	Number           rune = 35
	Dollar           rune = 36
	Percentage       rune = 37
	Ampersand        rune = 38
	SingleQuote      rune = 39
	LeftParenthesis  rune = 40
	RightParenthesis rune = 41
	Asterisk         rune = 42
	Plus             rune = 43
	Comma            rune = 44
	Hyphen           rune = 45
	Period           rune = 46
	ForwardSlash     rune = 47
	Zero             rune = 48
	One              rune = 49
	Two              rune = 50
	Three            rune = 51
	Four             rune = 52
	Five             rune = 53
	Six              rune = 54
	Seven            rune = 55
	Eight            rune = 56
	Nine             rune = 57
	Colon            rune = 58
	Semicolon        rune = 59
	LessThan         rune = 60
	Equals           rune = 61
	GreaterThan      rune = 62
	QuestionMark     rune = 63
	At               rune = 64
	UpperA           rune = 65
	UpperB           rune = 66
	UpperC           rune = 67
	UpperD           rune = 68
	UpperE           rune = 69
	UpperF           rune = 70
	UpperG           rune = 71
	UpperH           rune = 72
	UpperI           rune = 73
	UpperJ           rune = 74
	UpperK           rune = 75
	UpperL           rune = 76
	UpperM           rune = 77
	UpperN           rune = 78
	UpperO           rune = 79
	UpperP           rune = 80
	UpperQ           rune = 81
	UpperR           rune = 82
	UpperS           rune = 83
	UpperT           rune = 84
	UpperU           rune = 85
	UpperV           rune = 86
	UpperW           rune = 87
	UpperX           rune = 88
	UpperY           rune = 89
	UpperZ           rune = 90
	LeftBracket      rune = 91
	Backslash        rune = 92
	RightBracket     rune = 93
	Caret            rune = 94
	Underscore       rune = 95
	Grave            rune = 96
	LowerA           rune = 97
	LowerB           rune = 98
	LowerC           rune = 99
	LowerD           rune = 100
	LowerE           rune = 101
	LowerF           rune = 102
	LowerG           rune = 103
	LowerH           rune = 104
	LowerI           rune = 105
	LowerJ           rune = 106
	LowerK           rune = 107
	LowerL           rune = 108
	LowerM           rune = 109
	LowerN           rune = 110
	LowerO           rune = 111
	LowerP           rune = 112
	LowerQ           rune = 113
	LowerR           rune = 114
	LowerS           rune = 115
	LowerT           rune = 116
	LowerU           rune = 117
	LowerV           rune = 118
	LowerW           rune = 119
	LowerX           rune = 120
	LowerY           rune = 121
	LowerZ           rune = 122
	LeftBrace        rune = 123
	VerticalBar      rune = 124
	RightBrace       rune = 125
	Tilde            rune = 126
	Delete           rune = 127
)
