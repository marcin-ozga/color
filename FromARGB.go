package color

// Creates a new Color structure by using the specified sRGB alpha channel and color channel values.
//
// Parameters
//	alpha byte - The alpha channel, A, of the new color. A value between 0 and 255.
//	red   byte - The red channel, R, of the new color. A value between 0 and 255.
//	green byte - The green channel, G, of the new color. A value between 0 and 255.
//	blue  byte - The blue channel, B, of the new color. A value between 0 and 255.
//
// Returns
//	Color - A Color structure with the specified values.
//
// Remarks
//	The alpha channel of a color determines the amount of transparency of the color. An alpha value of 255 indicates that the color is completely opaque, and a value of 0 indicates that the color is completely transparent.
func FromARGB(alpha byte, red byte, green byte, blue byte) Color {
	return Color(uint32(alpha)<<alphaOffset | uint32(red)<<redOffset | uint32(green)<<greenOffset | uint32(blue)<<blueOffset)
}
