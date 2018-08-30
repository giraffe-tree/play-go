package main

import (
	"image/color"
	"math/rand"
	"time"
	"os"
	"net/http"
	"io"
	"image/gif"
	"image"
	"math"
	"log"
	"fmt"
)

var palette = []color.Color{color.White, color.RGBA{0xff, 0x00, 0x00, 0x00}, color.RGBA{0x00, 0xff, 0x00, 0x00}, color.RGBA{0x00, 0x00, 0xff, 0x00}}

const whiteIndex = 0
const blackIndex = 1

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)

}
func lissajous(out io.Writer) {
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
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			var x1 = uint8(rand.Uint32() % 3)+1
			fmt.Println(x1)
			x := math.Sin(t)
			y := math.Cos(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 2)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}
