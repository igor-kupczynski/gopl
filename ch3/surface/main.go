package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handleSurface)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSurface(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	drawSurface(w)
}

func drawSurface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ah := corner(i+1, j)
			bx, by, bh := corner(i, j)
			cx, cy, ch := corner(i, j+1)
			dx, dy, dh := corner(i+1, j+1)
			if numbers(ax, ay, ah, bx, by, bh, cx, cy, ch, dx, dy, dh) {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
					"style='fill: %s' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, color(ah, bh, ch, dh))
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute drawSurface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func numbers(xs ...float64) bool {
	for _, x := range xs {
		if math.IsNaN(x) {
			return false
		}
		if math.IsInf(x, 0) {
			return false
		}

	}
	return true
}

const (
	minBoost = 5
	maxBoost = 2
)

func color(h ...float64) string {
	var red, green, blue byte
	var max, min float64
	for _, x := range h {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	max *= maxBoost
	if max > 1.0 {
		max = 1.0
	}
	min *= minBoost
	if min < -1.0 {
		min = -1.0
	}

	maxDelta := byte(max * 255)
	minDelta := byte(-min * 255)

	red = 255 - minDelta/2
	green = 255 - (minDelta+maxDelta)/2
	blue = 255 - maxDelta/2

	return fmt.Sprintf("#%02x%02x%02x", red, green, blue)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
