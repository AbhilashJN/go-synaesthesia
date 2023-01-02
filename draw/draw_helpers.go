package draw

import (
	"go-synaesthesia/color"
	"go-synaesthesia/fft"
	"math"
	"math/cmplx"
	"math/rand"

	"github.com/fogleman/gg"
)

type coord struct {
	x int
	y int
}

func generateStartCoords(h, w int) []coord {
	sc := []coord{{x: 0, y: 0}}
	jumpSize := 55
	xstart := 0
	nextxstart := 0
	for y := 100; y < h-jumpSize; y += jumpSize {
		for x := xstart; x < w-jumpSize; x += jumpSize {
			sc = append(sc, coord{x, y})
		}
		xstart, nextxstart = nextxstart, xstart
	}
	return sc
}

func checkBounds(dc *gg.Context, x, y int) bool {
	if x >= 0 &&
		x < dc.Width() &&
		y >= 0 &&
		y < dc.Height() {
		return true
	}
	return false
}

func drawField(fftPhases [][]float64, outPath string) {
	dc := gg.NewContext(len(fftPhases), len(fftPhases[0]))
	for col_idx := 0; col_idx < len(fftPhases); col_idx += 200 {
		for row_idx := 0; row_idx < len(fftPhases[0]); row_idx += 200 {
			row_val := fftPhases[col_idx][row_idx]
			dc.SetRGB(0, 0, 1)
			dc.SetLineWidth(10)
			dc.DrawCircle(float64(col_idx), float64(row_idx), row_val)
			dc.Stroke()
		}
	}
	dc.SavePNG(outPath)
}

func visualizePhases(fftCoeffs [][]complex128, outPath string) {
	cols := len(fftCoeffs)
	rows := len(fftCoeffs[0])
	normalizedVals := make([][]float64, cols)
	for i := 0; i < cols; i++ {
		normalizedVals[i] = make([]float64, rows)
	}

	for col_idx, col_vals := range fftCoeffs {
		for row_idx, row_val := range col_vals {
			normalizedVals[col_idx][row_idx] = (cmplx.Phase(row_val) + math.Pi) / (2 * math.Pi)
		}
	}
	drawField(normalizedVals, outPath)
}

func drawLine(dc *gg.Context, fftCoeffs [][]complex128, numSteps, x, y int, lineWidth, stepSize float64, palette color.ColorPalette, alpha float64) {
	currentX, currentY := float64(x), float64(y)
	for i := 0; i < numSteps; i++ {
		nextX, nextY := int(currentX), int(currentY)
		if !checkBounds(dc, nextX, nextY) {
			break
		}
		if nextY == dc.Height()/2 {
			break
		}

		fftMag := fft.GetMagnitude(fftCoeffs[nextX][nextY])
		fftPhase := cmplx.Phase(fftCoeffs[nextX][nextY])

		xstep := stepSize
		ystep := fft.GetPowerToDb(fftMag) * 5
		lineWidth := math.Floor((fftPhase/math.Pi + 1) * lineWidth)

		currentX += xstep
		currentY += ystep
		dc.SetLineWidth(lineWidth)
		dc.LineTo(currentX, currentY)
		currentY = float64(y)
	}
	dc.Stroke()
}

func DrawLines(dc *gg.Context, fftCoeffs [][]complex128, palette color.ColorPalette) {
	var stepSize, baseLineWidth, lineWidth, alpha float64
	startCoords := generateStartCoords(dc.Height(), dc.Width())
	translation := rand.Float64()*2 + 2
	if rand.Float64() > 0.5 {
		translation = -translation
	}
	baseLineWidth = rand.Float64()*4 + 2
	dc.SetLineCapRound()
	strokePattern := color.StrokePattern{
		Palette:     palette,
		TotalHeight: dc.Height(),
		TotalWidth:  dc.Width(),
		Alpha:       0,
		Dimming:     1,
	}

	for _, coord := range startCoords {
		numSteps := 2
		stepSize = 1

		strokePattern.Alpha = 175
		strokePattern.Dimming = 1
		oldColorDistrbution := strokePattern.Palette.GetColorAt
		strokePattern.Palette.GetColorAt = color.BlackColorDistribution
		dc.Translate(translation, translation)
		lineWidth = baseLineWidth * 1.1
		dc.SetStrokeStyle(&strokePattern)
		drawLine(
			dc,
			fftCoeffs,
			numSteps,
			coord.x,
			coord.y,
			lineWidth,
			stepSize,
			palette,
			alpha,
		)

		strokePattern.Alpha = 255
		strokePattern.Palette.GetColorAt = oldColorDistrbution
		dc.Translate(-translation, -translation)
		lineWidth = baseLineWidth
		dc.SetStrokeStyle(&strokePattern)
		drawLine(
			dc,
			fftCoeffs,
			numSteps,
			coord.x,
			coord.y,
			lineWidth,
			stepSize,
			palette,
			alpha,
		)
	}
}

func DrawBackground(dc *gg.Context, bgPrimary, bgSecondary color.HSBColor) {
	xstart, ystart := 0, 0
	xend, yend := dc.Width(), dc.Height()
	bgPattern := color.BgPattern{
		BgPrimary:   bgPrimary,
		BgSecondary: bgSecondary,
		TotalHeight: yend,
	}

	for x := xstart; x <= xend; x++ {
		for y := ystart; y <= yend; y++ {
			col := bgPattern.ColorAt(x, y)
			dc.SetColor(col)
			dc.SetPixel(x, y)
		}
	}
	dc.Stroke()

	for x := xstart; x <= xend; x += 50 {
		for y := ystart; y <= yend; y += 50 {
			dc.SetLineWidth(0.4)
			dc.SetRGB(0.1, 0.1, 0.1)
			dc.DrawPoint(float64(x), float64(y), 0.6)
		}
	}
	dc.Stroke()
}

func DrawBorders(dc *gg.Context, bgPrimary color.HSBColor) {
	bgPrimaryDark := bgPrimary
	bgPrimaryDark.B -= 20
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.SetLineWidth(100)
	dc.SetRGB(color.HSBColorToRGB(bgPrimary))
	dc.Stroke()
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.SetLineWidth(75)
	dc.SetRGB(color.HSBColorToRGB(bgPrimaryDark))
	dc.Stroke()
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.SetLineWidth(50)
	dc.SetHexColor("#ffffff")
	dc.Stroke()
}
