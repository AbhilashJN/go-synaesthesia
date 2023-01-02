package main

import (
	"flag"
	"go-synaesthesia/color"
	"go-synaesthesia/draw"
	"go-synaesthesia/fft"
	"log"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

func main() {
	inPath := flag.String("i", "", "Path to .wav file")
	outPath := flag.String("o", "", "Path to output .png file")
	flag.Parse()
	if *inPath == "" {
		log.Fatal("input path is required")
	}
	if *outPath == "" {
		log.Fatal("output path is required")
	}
	fftCoeffs := fft.GetSTFTCoeffs(*inPath, 2048, 2048, 2048)

	maxWidth := 1600
	maxHeight := 1024
	if len(fftCoeffs) < maxWidth {
		maxWidth = len(fftCoeffs)
	}
	if len(fftCoeffs[0])/2 < maxHeight {
		maxHeight = len(fftCoeffs[0]) / 2
	}
	xstart := 0
	if len(fftCoeffs) > maxWidth {
		diff := len(fftCoeffs) - maxWidth
		xstart += diff / 2
	}

	fftCoeffsHalf := make([][]complex128, maxWidth)
	for i := xstart; i < xstart+maxWidth; i++ {
		fftCoeffsHalf[i-xstart] = make([]complex128, maxHeight)
		for j := 0; j < maxHeight; j++ {
			fftCoeffsHalf[i-xstart][j] = fftCoeffs[i][j]
		}
	}

	rand.Seed(time.Now().UnixNano())
	palette := color.Palettes[rand.Intn(len(color.Palettes))]
	dc := gg.NewContext(len(fftCoeffsHalf), len(fftCoeffsHalf[0]))
	draw.DrawBackground(dc, palette.BgPrimary, palette.BgSecondary)

	rotation := rand.Float64()/2 - 0.25
	dc.RotateAbout(rotation, float64(dc.Width())/2, float64(dc.Height())/2)
	draw.DrawLines(dc, fftCoeffsHalf, palette)
	dc.RotateAbout(-rotation, float64(dc.Width())/2, float64(dc.Height())/2)
	draw.DrawBorders(dc, palette.BgPrimary)
	err := dc.SavePNG(*outPath)
	if err != nil {
		log.Fatalf("error saving to png file %s: %v", *outPath, err)
	}
}
