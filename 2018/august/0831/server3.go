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
	"sync"
)

var mu1 sync.Mutex
var count1 int
var palette2 = []color.Color{color.White, color.RGBA{0xff, 0x00, 0x00, 0x00}, color.RGBA{0x00, 0xff, 0x00, 0x00}, color.RGBA{0x00, 0x00, 0xff, 0x00}}

func main() {

	http.HandleFunc("/", handler3)
	http.HandleFunc("/count", counter3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler3(w http.ResponseWriter, r *http.Request) {
	mu1.Lock()
	count1++
	mu1.Unlock()
	lissajous2(w)
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
	}
	fmt.Fprintf(w, "HOST = %q \n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}

func counter3(w http.ResponseWriter, r *http.Request) {
	mu1.Lock()
	fmt.Fprintf(w, "count %d\n", count1)
	mu1.Unlock()
}

/**
test
*/
func lissajous2(out io.Writer) {
	const (
		cycles  = 5     // x 振荡器变化的个数
		res     = 0.001 // 角度分辨率
		size    = 100
		nframes = 20
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0
	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			var x1 = uint8(rand.Uint32()%3) + 1
			x := math.Sin(t)
			y := math.Cos(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), x1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}
