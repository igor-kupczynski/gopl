package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var (
	visits int
	mux    sync.Mutex
)

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/__stats__", stats)
	http.HandleFunc("/lissa", lissa)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	visits++
	mux.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%s]: %v\n", k, v)
	}
	if err := r.ParseForm(); err != nil {
		log.Println("echo: %v", err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%s]: %v\n", k, v)
	}
}

func stats(w http.ResponseWriter, _ *http.Request) {
	mux.Lock()
	c := visits
	mux.Unlock()
	fmt.Fprintf(w, "Visits so far: %d", c)
}

var palette = []color.Color{
	color.White,
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // red
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // green
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // blue
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF},
}

func lissa(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("lissajous: %v", err)
	}
	cycles := 5.0
	cyclesStr := r.Form.Get("cycles")
	if cyclesStr != "" {
		i, err := strconv.Atoi(cyclesStr)
		if err == nil {
			cycles = float64(i)
		}
	}
	lissajous(w, cycles, palette)
}

func lissajous(out io.Writer, cycles float64, palette []color.Color) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	cls := 0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, size*2+1, size*2+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8((cls%(len(palette)-1))+1))
			cls++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	if err := gif.EncodeAll(out, &anim); err != nil {
		log.Fatalf("lissjous: %v", err)
	}
}
