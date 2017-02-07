package ColorAnsi

import "fmt"

const (
	ansiReset        = 0
	stringColorFormat = "\x1b[%d;%d;%dm%v\x1b[0m"
)

const (
	ansiIntensityOn  = 1
	ansiIntensityOff = 21
	ansiUnderlineOn  = 4
	ansiUnderlineOff = 24
	ansiBlinkOn      = 5
	ansiBlinkOff     = 25
)

const (
	ansiForegroundBlack   = 30
	ansiForegroundRed     = 31
	ansiForegroundGreen   = 32
	ansiForegroundYellow  = 33
	ansiForegroundBlue    = 34
	ansiForegroundMagenta = 35
	ansiForegroundCyan    = 36
	ansiForegroundWhite   = 37
	ansiForegroundDefault = 39
)

const (
	ansiBackgroundBlack   = 40
	ansiBackgroundRed     = 41
	ansiBackgroundGreen   = 42
	ansiBackgroundYellow  = 43
	ansiBackgroundBlue    = 44
	ansiBackgroundMagenta = 45
	ansiBackgroundCyan    = 46
	ansiBackgroundWhite   = 47
	ansiBackgroundDefault = 49
)

const (
	ansiLightForegroundGray    = 90
	ansiLightForegroundRed     = 91
	ansiLightForegroundGreen   = 92
	ansiLightForegroundYellow  = 93
	ansiLightForegroundBlue    = 94
	ansiLightForegroundMagenta = 95
	ansiLightForegroundCyan    = 96
	ansiLightForegroundWhite   = 97
)

const (
	ansiLightBackgroundGray    = 100
	ansiLightBackgroundRed     = 101
	ansiLightBackgroundGreen   = 102
	ansiLightBackgroundYellow  = 103
	ansiLightBackgroundBlue    = 104
	ansiLightBackgroundMagenta = 105
	ansiLightBackgroundCyan    = 106
	ansiLightBackgroundWhite   = 107
)

type Color struct {
	style int
	foreground int
	background int
}

//Clear
func (c *Color)  ClearForeground()  {
	c.foreground = ansiReset
}

func (c *Color)  ClearBackground()  {
	c.background = ansiReset
}

func (c *Color) ClearStyle()  {
	c.style = ansiReset
}

//Default
func (c *Color) SetForegroundDefault()  {
	c.foreground = ansiForegroundDefault
}

func (c *Color) SetBackgroundDefault()  {
	c.background = ansiBackgroundDefault
}

//Style
func (c *Color) SetIntensityOn ()  {
	c.style = ansiIntensityOn
}

func (c *Color) SetIntensityOff ()  {
	c.style = ansiIntensityOff
}

func (c *Color) SetUnderlineOn ()  {
	c.style = ansiUnderlineOn
}

func (c *Color) SetUnderlineOff ()  {
	c.style = ansiUnderlineOff
}

func (c *Color) SetBlinkOn ()  {
	c.style = ansiBlinkOn
}

func (c *Color) SetBlinkOff ()  {
	c.style = ansiBlinkOff
}

//Black
func (c *Color) SetForegroundBlack()  {
	c.foreground = ansiForegroundBlack
}

func (c *Color) SetBackgroundBlack()  {
	c.background = ansiBackgroundBlack
}

//Red
func (c *Color) SetForegroundRed()  {
	c.foreground = ansiForegroundRed
}

func (c *Color) SetBackgroundRed()  {
	c.background = ansiBackgroundRed
}

//Green
func (c *Color) SetForegroundGreen()  {
	c.foreground = ansiForegroundGreen
}

func (c *Color) SetBackgroundGreen()  {
	c.background = ansiBackgroundGreen
}

//Yellow
func (c *Color) SetForegroundYellow()  {
	c.foreground = ansiForegroundYellow
}

func (c *Color) SetBackgroundYellow()  {
	c.background = ansiBackgroundYellow
}

//Blue
func (c *Color) SetForegroundBlue()  {
	c.foreground = ansiForegroundBlue
}

func (c *Color) SetBackgroundBlue()  {
	c.background = ansiBackgroundBlue
}

//Magenta
func (c *Color) SetForegroundMagenta()  {
	c.foreground = ansiForegroundMagenta
}

func (c *Color) SetBackgroundMagenta()  {
	c.background = ansiBackgroundMagenta
}

//Cyan
func (c *Color) SetForegroundCyan()  {
	c.foreground = ansiForegroundCyan
}

func (c *Color) SetBackgroundCyan()  {
	c.background = ansiBackgroundCyan
}

//White
func (c *Color) SetForegroundWhite()  {
	c.foreground = ansiForegroundWhite
}

func (c *Color) SetBackgroundWhite()  {
	c.background = ansiBackgroundWhite
}

///////////////////////////////////////////////////////////////////////////////
// Light
///////////////////////////////////////////////////////////////////////////////

//LightGray
func (c *Color) SetForegroundLightGray()  {
	c.foreground = ansiLightForegroundGray
}

func (c *Color) SetBackgroundLightGray()  {
	c.background = ansiLightBackgroundGray
}

//LightRed
func (c *Color) SetForegroundLightRed()  {
	c.foreground = ansiLightForegroundRed
}

func (c *Color) SetBackgroundLightRed()  {
	c.background = ansiLightBackgroundRed
}

//LightGreen
func (c *Color) SetForegroundLightGreen()  {
	c.foreground = ansiLightForegroundGreen
}

func (c *Color) SetBackgroundLightGreen()  {
	c.background = ansiLightBackgroundGreen
}

//LightYellow
func (c *Color) SetForegroundLightYellow()  {
	c.foreground = ansiLightForegroundYellow
}

func (c *Color) SetBackgroundLightYellow()  {
	c.background = ansiLightBackgroundYellow
}

//LightBlue
func (c *Color) SetForegroundLightBlue()  {
	c.foreground = ansiLightForegroundBlue
}

func (c *Color) SetBackgroundLightBlue()  {
	c.background = ansiLightBackgroundBlue
}

//LightMagenta
func (c *Color) SetForegroundLightMagenta()  {
	c.foreground = ansiLightForegroundMagenta
}

func (c *Color) SetBackgroundLightMagenta()  {
	c.background = ansiLightBackgroundMagenta
}

//LightCyan
func (c *Color) SetForegroundLightCyan()  {
	c.foreground = ansiLightForegroundCyan
}

func (c *Color) SetBackgroundLightCyan()  {
	c.background = ansiLightBackgroundCyan
}

//LightWhite
func (c *Color) SetForegroundLightWhite()  {
	c.foreground = ansiLightForegroundWhite
}

func (c *Color) SetBackgroundLightWhite()  {
	c.background = ansiLightBackgroundWhite
}

func (c *Color) GetString(str string) string {
	return fmt.Sprintf(stringColorFormat, c.style,  c.foreground,  c.background, str)
}



