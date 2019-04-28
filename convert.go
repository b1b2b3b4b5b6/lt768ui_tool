package main

import (
	"goc/toolcom/errtool"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func convert2bin(path string) {
	file, err := os.Open(path)
	errtool.Errpanic(err)
	defer file.Close()
	img, _, err := image.Decode(file)
	errtool.Errpanic(err)
	bound := img.Bounds()
	raw := make([]color.Color, 0)

	log.Debug("x[%d] y[%d]", bound.Max.X, bound.Max.Y)
	for y := 0; y < bound.Max.Y; y++ {
		for x := 0; x < bound.Max.X; x++ {
			raw = append(raw, img.At(x, y))
		}
	}

	// rgb16 := make([]uint16, 0)

	// for _, val := range raw {
	// 	_r, _g, _b, _a := val.RGBA()
	// 	var r, g, b, a uint8
	// 	r = uint8(_r / 255)
	// 	g = uint8(_g / 255)
	// 	b = uint8(_b / 255)
	// 	a = uint8(_a)
	// 	log.Debug("[%d] [%d] [%d] [%d]", r, g, b, a)
	// 	if a != 255 {
	// 		rgb16 = append(rgb16, 0)
	// 		continue
	// 	}
	// 	r >>= 3
	// 	g >>= 2
	// 	b >>= 3
	// 	log.Debug("[%d] [%d] [%d]", r, g, b)
	// 	var rgb uint16
	// 	rgb = uint16(r)
	// 	rgb <<= 6
	// 	rgb |= uint16(g)
	// 	rgb <<= 5
	// 	rgb |= uint16(b)

	// 	rgb16 = append(rgb16, rgb)
	// }
	log.Debug("raw[%+v]", raw)
	// log.Debug("rgb16[%+v]", rgb16)
}
