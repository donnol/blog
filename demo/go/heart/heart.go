package heart

import (
	"image"
	"image/color"
	"image/draw"
	"os"
)

func Draw() {
	rect := image.Rect(100, 100, 100, 100)
	rgb := image.NewNRGBA(rect)
	rgb.Set(10, 10, color.Black)

	draw.Draw(rgb, rect, rgb, image.Pt(20, 20), draw.Over)
	if err := os.WriteFile("heart.png", []byte(rgb.Rect.String()), os.ModePerm); err != nil {
		panic(err)
	}
}
