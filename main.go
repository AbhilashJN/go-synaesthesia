package main

import (
	"flag"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/wav"
)

func getFFTCoeffs(filePath string, ts int) []complex128 {
	fd, err := os.Open(filePath)
	if err != nil {
		log.Fatal("error opening file")
	}
	wavStruct, err := wav.New(fd)
	rand.Seed(int64(wavStruct.Samples))
	if err != nil {
		log.Fatal("error reading wav")
	}
	totalSamples := int(math.Min(float64(ts), float64(wavStruct.Samples)))

	chunkSize := wavStruct.Samples / totalSamples
	if chunkSize == 0 {
		chunkSize = 1
	}
	waveSamplesFloat64 := make([]float64, totalSamples)

	for i := 0; i < totalSamples; i++ {
		wavSamples, err := wavStruct.ReadFloats(chunkSize)
		if err != nil {
			log.Fatal("error reading samples")
		}
		waveSamplesFloat64 = append(waveSamplesFloat64, float64(wavSamples[0]))
	}
	return fft.FFTReal(waveSamplesFloat64)[:totalSamples]
}

func drawSection(dc *gg.Context, section []complex128, cx, cy float64, rotationFactor float64) {
	dc.MoveTo(cx, cy)
	rand.Seed(time.Now().UnixNano())
	dc.RotateAbout(-math.Pi/rotationFactor, cx, cy)
	currentX, currentY := cx, cy
	for _, num := range section {
		x, y := real(num), imag(num)
		x, y = normalize(x), normalize(y)
		dc.MoveTo(currentX, currentY)
		dc.LineTo(currentX+x, currentY+y)
		dc.RotateAbout(-math.Pi/rotationFactor, cx+x, cy+y)
		rcolor := rand.Float64()
		gcolor := rand.Float64()
		bcolor := rand.Float64()
		dc.SetRGBA(rcolor, gcolor, bcolor, 0.8)
		dc.SetLineWidth(10)
		dc.SetLineCap(gg.LineCapRound)
		dc.Stroke()
		currentX, currentY = currentX+x, currentY+y
	}

	// dc.SetRGBA(math.Abs(math.Sin(currentX)), math.Abs(math.Cos(currentX)), math.Abs(math.Sin(currentX)), 1)
	// dc.SetLineWidth(10)
	// dc.SetLineCap(gg.LineCapRound)
	// dc.Stroke()
}

func normalize(num float64) float64 {
	return math.Floor(100 * (2/(1+(math.Pow(math.E, -10*num))) - 1))
}

func main() {
	inPath := flag.String("i", "track.wav", "Path to .wav file")
	outPath := flag.String("o", "out.png", "Path to output .png file")
	flag.Parse()
	fftCoeffs := getFFTCoeffs(*inPath, 1000000)
	var width, height = 2000, 2000
	widthFloat, heightFloat := float64(width), float64(height)
	cx, cy := widthFloat/2, heightFloat/2
	dc := gg.NewContext(width, height)
	dc.MoveTo(0, 0)
	dc.DrawRectangle(0, 0, widthFloat, heightFloat)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	rotationFactor := float64((rand.Intn(2) + 1) * 2)

	for i := 0; i < 1000; i += 1 {
		begin := i * 100
		section := fftCoeffs[begin : begin+50]
		drawSection(dc, section, cx, cy, rotationFactor)
	}

	dc.SavePNG(*outPath)
}
