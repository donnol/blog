package heart

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"os"
)

func Draw() {
	rect := image.Rect(0, 0, 100, 100) // 对角线坐标
	rgb := image.NewNRGBA(rect)

	// 逐个点涂色
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			rgb.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, rgb); err != nil {
		panic(err)
	}
	if err := os.WriteFile("heart.png", buf.Bytes(), os.ModePerm); err != nil {
		panic(err)
	}
}
