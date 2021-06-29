package tool

import (
	"image"
	"image/color"
)

func DrawVerticalLine(img *image.RGBA, x int, y1 int, y2 int, lineWidth int, clr color.Color) {
	if x < img.Rect.Min.X || x > img.Rect.Max.X {
		return
	}
	if y1 < img.Rect.Min.Y {
		y1 = img.Rect.Min.Y
	}
	if y2 > img.Rect.Max.Y {
		y2 = img.Rect.Max.Y 
	}
	if img.Rect.Max.X - x + 1 < lineWidth {
		lineWidth = img.Rect.Max.X - x + 1
	}
	for y := y1; y <= y2; y++ {
		for line1px := 0; line1px < lineWidth; line1px++ {
			img.Set(x + line1px, y, clr)
		}
	}
}

func DrawHorisontalLine(img *image.RGBA, x1 int, x2 int, y int, lineWidth int, clr color.Color) {
	if y < img.Rect.Min.Y || y > img.Rect.Max.Y {
		return
	}
	if x1 < img.Rect.Min.X {
		x1 = img.Rect.Min.X
	}
	if x2 > img.Rect.Max.X {
		x2 = img.Rect.Max.X 
	}
	if img.Rect.Max.Y - y + 1 < lineWidth {
		lineWidth = img.Rect.Max.Y - y + 1
	}
	for x := x1; x <= x2; x++ {
		for line1px := 0; line1px < lineWidth; line1px++ {
			img.Set(x, y + line1px, clr)
		}
	}
}