package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handleDraw)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleDraw(w http.ResponseWriter, r *http.Request) {
	sampleRate, err := strconv.Atoi(r.FormValue("rate"))
	if err != nil {
		sampleRate = 1
	}
	draw(w, sampleRate)
}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func draw(out io.Writer, sampleRate int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			// Image point (px, py) represents complex value z.
			img.Set(px, py, getColor(supersample(sampleRate, x, y, mandelbrot)))
		}
	}
	_ = png.Encode(out, img) // NOTE: ignoring errors
}

func supersample(sampleRate int, x, y float64, f func(float64, float64) byte) byte {
	if sampleRate == 1 {
		// No supersampling
		return f(x, y)
	}

	dx := float64(xmax-xmin) / float64(width*sampleRate)
	dy := float64(ymax-ymin) / float64(height*sampleRate)

	var total int
	for yi := 0; yi < sampleRate; yi++ {
		sy := y + dy*float64(sampleRate/2-yi)
		for xi := 0; xi < sampleRate; xi++ {
			sx := x + dx*float64(sampleRate/2-xi)
			total += int(f(sx, sy))
		}
	}

	return byte(total / (sampleRate * sampleRate))

}

func getColor(n byte) color.Color {
	const contrast = 15
	c := 255 - contrast*n
	if c < 0 {
		c = 0
	}
	return color.Gray{Y: c}
}

func mandelbrot(x, y float64) byte {
	z := complex(x, y)
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return 255
}
