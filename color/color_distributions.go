package color

type ColorDistributor func(x, y, h, w int, colors []HSBColor) HSBColor

func Lerp4Corners(x, y, h, w int, colors []HSBColor) HSBColor {
	lerpH1 := HSBLerp(colors[0], colors[1], x, w)
	lerpH2 := HSBLerp(colors[2], colors[3], x, w)
	return HSBLerp(lerpH1, lerpH2, y, h)
}

func GridDistribution(x, y, h, w int, colors []HSBColor) HSBColor {
	xrem := x % 400
	yrem := y % 400
	colorIdx := 0
	if xrem >= 200 {
		colorIdx += 1
	}
	if yrem >= 200 {
		colorIdx += 2
	}
	return colors[colorIdx]
}
func ThreeColorDistribution(x, y, h, w int, colors []HSBColor) HSBColor {
	idx := 0
	switch {
	case y%300 < 100:
		idx = 0
	case y%300 < 200:
		idx = 1
	default:
		idx = 2
	}
	return colors[idx]
}

func BlackColorDistribution(x, y, h, w int, colors []HSBColor) HSBColor {
	return newHSBColor(0, 0, 0)
}
