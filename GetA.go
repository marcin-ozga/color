package color

// Gets the alpha channel value of the color.
//
// Returns
//	The alpha channel value of the Color structure, a value between 0 and 255.
func (color Color) GetAlpha() byte {
	return byte(color >> alphaOffset)
}

// Gets the red channel value of the color.
//
// Returns
//	The red channel value of the Color structure, a value between 0 and 255.
func (color Color) GetRed() byte {
	return byte(color >> redOffset)
}

// Gets the green channel value of the color.
//
// Returns
//	The green channel value of the Color structure, a value between 0 and 255.
func (color Color) GetGreen() byte {
	return byte(color >> greenOffset)
}

// Gets the blue channel value of the color.
//
// Returns
//	The blue channel value of the Color structure, a value between 0 and 255.
func (color Color) GetBlue() byte {
	return byte(color >> blueOffset)
}

// Sets the alpha channel value of the color.
//
// Parameters
//	alpha byte - The alpha channel value of the Color structure, a value between 0 and 255.
func (color Color) SetAlpha(alpha byte) {
	color = color & ^Color(alphaMask) | Color(alpha)<<alphaOffset
}

// Sets the red channel value of the color.
//
// Parameters
//	red byte - The red channel value of the Color structure, a value between 0 and 255.
func (color Color) SetRed(red byte) {
	color = color & ^Color(redMask) | Color(red)<<redOffset
}

// Sets the green channel value of the color.
//
// Parameters
//	green byte - The green channel value of the Color structure, a value between 0 and 255.
func (color Color) SetGreen(green byte) {
	color = color & ^Color(greenMask) | Color(green)<<greenOffset
}

// Sets the blue channel value of the color.
//
// Parameters
//	blue byte - The blue channel value of the Color structure, a value between 0 and 255.
func (color Color) SetBlue(blue byte) {
	color = color & ^Color(blueMask) | Color(blue)<<blueOffset
}
