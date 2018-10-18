// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	// const (
	// 	cycles  = 5     // number of complete x oscillator revolutions
	// 	res     = 0.001 // angular resolution
	// 	size    = 100   // image canvas covers [-size..+size]
	// 	nframes = 64    // number of animation frames
	// 	delay   = 8     // delay between frames in 10ms units
	// )
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// for k, v := range r.Header {
	// 	fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// }
	// fmt.Fprintf(w, "Host = %q\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// for k, v := range r.Form {
	// 	fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	// }

	var (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	if parsed, err := strconv.Atoi(r.Form.Get("cycles")); err == nil {
		cycles = parsed
	}
	if parsed, err := strconv.ParseFloat(r.Form.Get("res"), 64); err == nil {
		res = parsed
	}
	if parsed, err := strconv.Atoi(r.Form.Get("size")); err == nil {
		size = parsed
	}
	if parsed, err := strconv.Atoi(r.Form.Get("nframes")); err == nil {
		nframes = parsed
	}
	if parsed, err := strconv.Atoi(r.Form.Get("delay")); err == nil {
		delay = parsed
	}

	// fmt.Fprintf(w, "cycles = %d\n", cycles)
	// fmt.Fprintf(w, "res = %f\n", res)
	// fmt.Fprintf(w, "size = %d\n", size)
	// fmt.Fprintf(w, "nframes = %d\n", nframes)
	// fmt.Fprintf(w, "delay = %d\n", delay)

	lissajous(w, cycles, res, size, nframes, delay)
}

//!-handler
