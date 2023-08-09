package color

import "fmt"

func (color Color) String() string {
	return fmt.Sprintf("#%08X", uint32(color))
}
