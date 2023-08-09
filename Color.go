package color

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type Color uint32

const (
	alphaOffset = int(24)
	redOffset   = int(16)
	greenOffset = int(8)
	blueOffset  = int(0)
)

const (
	alphaMask = Color(0xFF000000)
	redMask   = Color(0x00FF0000)
	greenMask = Color(0x0000FF00)
	blueMask  = Color(0x000000FF)
	rgbMask   = Color(0x00FFFFFF)
)

func (color Color) MarshalJSON() ([]byte, error) {
	return json.Marshal(color.String())
}

func (color *Color) UnmarshalJSON(data []byte) error {
	colorString := strings.Trim(string(data), "\"")
	if colorString == "" {
		*color = 0
		return nil
	}
	if colorString[0] != '#' {
		return errors.New("bad syntax")
	}
	colorUint64, err := strconv.ParseUint(colorString[1:], 16, 64)
	if err != nil {
		return err
	}
	*color = Color(colorUint64)
	return nil
}
