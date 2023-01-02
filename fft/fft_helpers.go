package fft

import (
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"os"

	"github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/wav"
	"github.com/r9y9/gossp/stft"
	"github.com/r9y9/gossp/window"
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

func GetSTFTCoeffs(filePath string, frameSize, windowSize, hopSize int) [][]complex128 {
	fd, err := os.Open(filePath)
	if err != nil {
		log.Fatal("error opening file ", filePath)
	}
	wavStruct, err := wav.New(fd)
	if err != nil {
		log.Fatal("error reading wav")
	}
	waveSamplesFloat64 := make([]float64, wavStruct.Samples)
	waveSamplesFloat32, err := wavStruct.ReadFloats(wavStruct.Samples)
	if err != nil {
		log.Fatal("error reading floats")
	}
	for i, s := range waveSamplesFloat32 {
		waveSamplesFloat64[i] = float64(s)
	}

	s := &stft.STFT{
		FrameShift: hopSize,
		FrameLen:   frameSize,
		Window:     window.CreateHanning(windowSize),
	}
	return s.STFT(waveSamplesFloat64)
}

func coeffsToAbsSqr(fftCoeffs [][]complex128) [][]float64 {
	fftAbsSqr := make([][]float64, len(fftCoeffs))

	for i, frame := range fftCoeffs {
		fftAbsSqr[i] = make([]float64, len(frame))
		for j, val := range frame {
			fftAbsSqr[i][j] = cmplx.Abs(val) * cmplx.Abs(val)
		}
	}

	return fftAbsSqr
}

func powerToDb(fftAbsSqr [][]float64) [][]float64 {
	fftDb := make([][]float64, len(fftAbsSqr))
	ref := 10 * math.Log10(1)
	for i, frame := range fftAbsSqr {
		fftDb[i] = make([]float64, len(frame))
		for j, val := range frame {
			fftDb[i][j] = 10*math.Log10(val) - ref
		}
	}

	return fftDb
}

func GetMagnitude(val complex128) float64 {
	return cmplx.Abs(val) * cmplx.Abs(val)
}

func GetPowerToDb(val float64) float64 {
	ref := 10 * math.Log10(1)
	return 10*math.Log10(val) - ref
}
