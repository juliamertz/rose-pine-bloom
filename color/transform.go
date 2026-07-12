package color

const (
	maxLightness = 100
	maxSteps     = 10
)

func Lighten(color *Color, steps uint8) Color {
	stepSize := (maxLightness - color.HSL.L) / maxSteps
	l := color.HSL.L + steps*stepSize
	if l > maxLightness || l < color.HSL.L {
		l = maxLightness
	}

	return ColorFromHSL(HSL{
		H: color.HSL.H,
		S: color.HSL.S,
		L: l,
	})
}

func Darken(color *Color, steps uint8) Color {
	stepSize := color.HSL.L / maxSteps
	sub := steps * stepSize
	l := color.HSL.L
	if sub < l {
		l -= sub
	} else {
		l = 0
	}

	return ColorFromHSL(HSL{
		H: color.HSL.H,
		S: color.HSL.S,
		L: l,
	})
}
