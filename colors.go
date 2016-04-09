package cc

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

var enabled = false

var osStdout *os.File = os.Stdout
var osStderr *os.File = os.Stderr

func doFormat(s string, code ...int) string {
	f := make([]string, len(code))
	for i, c := range code {
		f[i] = strconv.Itoa(c)
	}

	if s[len(s)-1] == '\n' {
		return fmt.Sprintf("\x1b[%sm%s\x1b[0m\n", strings.Join(f, ";"), s[:len(s)-1])
	}

	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", strings.Join(f, ";"), s)
}

func noFormat(s string, code ...int) string {
	return s
}

var fmfn = noFormat

// Enabled returns true if coloring is enabled.
func Enabled() bool {
	return enabled
}

// Enable enables coloring.
func Enable() {
	if enabled {
		return
	}

	os.Stdout = colorable.NewColorableStdout().(*os.File)
	os.Stderr = colorable.NewColorableStderr().(*os.File)

	fmfn = doFormat

	enabled = true
}

// Disable disables coloring.
func Disable() {
	if !enabled {
		return
	}

	os.Stdout = osStdout
	os.Stderr = osStderr

	fmfn = noFormat

	enabled = false
}

// Dark black foreground color
func BlackL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgBlack) }

// Dark red foreground color
func RedL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgRed) }

// Dark green foreground color
func GreenL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgGreen) }

// Dark yellow foreground color
func YellowL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgYellow) }

// Dark blue foreground color
func BlueL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgBlue) }

// Dark magenta foreground color
func MagentaL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgMagenta) }

// Dark cyan foreground color
func CyanL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgCyan) }

// Dark white foreground color
func WhiteL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgWhite) }

// Hi-intensity black foreground color
func BlackH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghBlack) }

// Hi-intensity red foreground color
func RedH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghRed) }

// Hi-intensity green foreground color
func GreenH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghGreen) }

// Hi-intensity yellow foreground color
func YellowH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghYellow) }

// Hi-intensity blue foreground color
func BlueH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghBlue) }

// Hi-intensity magenta foreground color
func MagentaH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghMagenta) }

// Hi-intensity cyan foreground color
func CyanH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghCyan) }

// Hi-intensity white foreground color
func WhiteH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghWhite) }

// Dark black background color
func BgBlackL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgBlack) }

// Dark red background color
func BgRedL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgRed) }

// Dark green background color
func BgGreenL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgGreen) }

// Dark yellow background color
func BgYellowL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgYellow) }

// Dark blue background color
func BgBlueL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgBlue) }

// Dark magenta background color
func BgMagentaL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgMagenta) }

// Dark cyan background color
func BgCyanL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgCyan) }

// Dark white background color
func BgWhiteL(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgWhite) }

// Hi-intensity black background color
func BgBlackH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghBlack) }

// Hi-intensity red background color
func BgRedH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghRed) }

// Hi-intensity green background color
func BgGreenH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghGreen) }

// Hi-intensity yellow background color
func BgYellowH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghYellow) }

// Hi-intensity bluebackground color
func BgBlueH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghBlue) }

// Hi-intensity magenta background color
func BgMagentaH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghMagenta) }

// Hi-intensity cyan background color
func BgCyanH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghCyan) }

// Hi-intensity white background color
func BgWhiteH(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghWhite) }

// Black foreground color
func Black(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgBlack) }

// Red foreground color
func Red(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghRed) }

// Green foreground color
func Green(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghGreen) }

// Yellow foreground color
func Yellow(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghYellow) }

// Blue foreground color
func Blue(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghBlue) }

// Magenta foreground color
func Magenta(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghMagenta) }

// Cyan foreground color
func Cyan(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fghCyan) }

// White foreground color
func White(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), fgWhite) }

// Black background color
func BgBlack(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgBlack) }

// Red background color
func BgRed(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghRed) }

// Green background color
func BgGreen(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghGreen) }

// Yellow background color
func BgYellow(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghYellow) }

// Blue background color
func BgBlue(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghBlue) }

// Magenta background color
func BgMagenta(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghMagenta) }

// Cyan background color
func BgCyan(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bghCyan) }

// White background color
func BgWhite(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), bgWhite) }

// Bold style
func Bold(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), styleBold) }

// Faint style
func Faint(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), styleFaint) }

// Italic style
func Italic(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), styleItalic) }

// Underline style
func Underline(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), styleUnderline) }

// Blink-Slow style
func BlinkSlow(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), styleBlinkSlow) }

// Blink-Rapid style
func BlinkRapid(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), styleBlinkRapid) }

// Reverse-Video style
func ReverseVideo(s string, a ...interface{}) string {
	return fmfn(fmt.Sprintf(s, a...), styleReverseVideo)
}

// Concealed style
func Concealed(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), styleConcealed) }

// Crossed-Out style
func CrossedOut(s string, a ...interface{}) string { return fmfn(fmt.Sprintf(s, a...), styleCrossedOut) }

func init() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		Enable()
		return
	}

	Disable()
}
