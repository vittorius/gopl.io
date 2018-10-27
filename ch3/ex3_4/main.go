// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells = 100         // number of grid cells
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	var (
		width  int64 = 600     // canvas width in pixels
		height int64 = 320     // canvas height in pixels
		fill         = "white" // fill
	)

	r.ParseForm()

	if parsed, err := strconv.Atoi(r.Form.Get("width")); err == nil {
		width = int64(parsed)
	}
	if parsed, err := strconv.Atoi(r.Form.Get("height")); err == nil {
		height = int64(parsed)
	}
	if parsed := r.Form.Get("fill"); parsed != "" {
		fill = parsed
	}

	// log.Printf("wigth: %d, height: %d, fill: %s", width, height, fill)

	writeSvg(w, width, height, fill)
}

func writeSvg(w io.Writer, width int64, height int64, fill string) {
	var (
		xyrange = 30.0                         // axis ranges (-xyrange..+xyrange)
		xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
		zscale  = float64(height) * 0.4        // pixels per z unit
	)

	// log.Printf("xyrange: %g, xyscale: %g, zscale: %g", xyrange, xyscale, zscale)

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", fill, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, xyrange, xyscale, zscale)
			bx, by := corner(i, j, width, height, xyrange, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, xyrange, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, xyrange, xyscale, zscale)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i int, j int, width int64, height int64, xyrange float64, xyscale float64, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
