package color

import "fmt"

const (
	FormatARGB = "ARGB"
	FormatRGB  = "RGB"
	FormatRGBA = "RGBA"
)

func (color Color) ToString(format string) string {
	switch format {
	case FormatARGB:
		return fmt.Sprintf("#%08X", uint32(color))
	case FormatRGB:
		return fmt.Sprintf("#%06X", uint32(color&rgbMask))
	case FormatRGBA:
		return fmt.Sprintf("#%08X", uint32((color&rgbMask)<<greenOffset|color>>alphaOffset))
	default:
		return fmt.Sprintf(format, uint32(color))
	}
}
