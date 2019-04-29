package main

import (
	"encoding/binary"
	"goc/toolcom/cfgtool"
	"goc/toolcom/errtool"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func convert2bin(path string) ([]byte, int, int) {

	nullColor, err := cfgtool.New("conf.json").TakeInt("NullColor")
	errtool.Errpanic(err)

	file, err := os.Open(path)
	errtool.Errpanic(err)
	defer file.Close()
	img, _, err := image.Decode(file)
	errtool.Errpanic(err)
	bound := img.Bounds()
	raw := make([]color.NRGBA, 0)

	for y := 0; y < bound.Max.Y; y++ {
		for x := 0; x < bound.Max.X; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y))
			nrgb := c.(color.NRGBA)
			raw = append(raw, nrgb)
		}
	}

	rgb16 := make([]uint16, 0)

	for _, val := range raw {
		r := val.R >> 3
		g := val.G >> 2
		b := val.B >> 3

		if val.A != 255 {
			rgb16 = append(rgb16, uint16(nullColor))
			continue
		}

		var rgb uint16
		rgb = uint16(r)
		rgb <<= 6
		rgb |= uint16(g)
		rgb <<= 5
		rgb |= uint16(b)
		rgb16 = append(rgb16, rgb)

	}

	endian, err := cfgtool.New("conf.json").TakeString("Endian")
	errtool.Errpanic(err)

	rgb8 := make([]byte, 0)
	for _, val := range rgb16 {
		b8 := make([]byte, 2)
		switch endian {
		case "BigEndian":
			binary.BigEndian.PutUint16(b8, val)
		case "LittleEndian":
			binary.LittleEndian.PutUint16(b8, val)
		default:
			log.Panic("unknown endian[%s]", endian)
		}
		rgb8 = append(rgb8, b8...)
	}
	return rgb8, bound.Max.X, bound.Max.Y
}
