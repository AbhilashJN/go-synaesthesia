package color

import (
	"image/color"
	"math"
)

type HSBColor struct {
	H int
	S int
	B int
}

func newHSBColor(h, s, b int) HSBColor {
	return HSBColor{
		H: h, S: s, B: b,
	}
}

func HSBLerp(c1, c2 HSBColor, param, max int) HSBColor {
	hgap := float64(c1.H - c2.H)
	sgap := float64(c1.S - c2.S)
	Bgap := float64(c1.B - c2.B)

	return newHSBColor(
		c1.H+int(hgap*float64(param)/float64(max)),
		c1.S+int(sgap*float64(param)/float64(max)),
		c1.B+int(Bgap*float64(param)/float64(max)),
	)
}

type ColorPalette struct {
	BgPrimary   HSBColor
	BgSecondary HSBColor
	Colors      []HSBColor
	GetColorAt  ColorDistributor
}

type BgPattern struct {
	BgPrimary   HSBColor
	BgSecondary HSBColor
	TotalHeight int
}

func (p *BgPattern) ColorAt(x, y int) color.Color {
	l := HSBLerp(p.BgSecondary, p.BgPrimary, y, p.TotalHeight)
	r, g, b := HSBColorToRGB(l)
	return color.NRGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: 255,
	}
}

type StrokePattern struct {
	Palette     ColorPalette
	TotalHeight int
	TotalWidth  int
	Alpha       int
	Dimming     float64
}

func (p *StrokePattern) ColorAt(x, y int) color.Color {
	lerpColor := p.Palette.GetColorAt(x, y, p.TotalHeight, p.TotalWidth, p.Palette.Colors)
	lerpColor.B = int(float64(lerpColor.B) * p.Dimming)
	r, g, b := HSBColorToRGB(lerpColor)
	return color.NRGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: uint8(p.Alpha),
	}
}

func HSBColorToRGB(c HSBColor) (float64, float64, float64) {
	return HSBToRGB(c.H, c.S, c.B)
}

func HSBToRGB(h, s, b int) (float64, float64, float64) {
	sf := float64(s) / 100
	bf := float64(b) / 100
	hf := float64(h)
	k := func(n float64) float64 {
		div := n + hf/60
		return math.Mod(div, 6)
	}
	f := func(n float64) float64 {
		min := k(n)
		if 4-k(n) < min {
			min = 4 - k(n)
		}
		if 1 < min {
			min = 1
		}
		return bf * (1 - sf*math.Max(0, min))
	}
	return 1 * f(5), 1 * f(3), 1 * f(1)
}
