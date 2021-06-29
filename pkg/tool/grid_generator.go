package tool

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	None = iota
	Numeric
	Alphabetic
)

type GridGeneratorInput struct {
	DPI                   float64
	VerticalCellsNum      int
	HorisontalCellsNum    int
	CellHeight            int
	CellWidth             int
	FrameWidth            int
	VerticalCellsMarker   int
	HorisontalCellsMarker int
	BackgroundColor       color.RGBA
	FrameColor            color.RGBA
	Font                  *truetype.Font
	FontSize              float64
	TextPaddingLeft       int
	TextPaddingTop        int
	FontColor             color.RGBA
	MarkerOX              Marker
	MarkerOY              Marker
	Delimeter             string
}

func (g *GridGeneratorInput) Valide() error {
	checkVerticalMarker := g.VerticalCellsMarker == Numeric ||
		g.VerticalCellsMarker == Alphabetic
	if !checkVerticalMarker {
		return fmt.Errorf("wrong type of vertical marker")
	}
	checkHorizontalmarker := g.HorisontalCellsMarker == Numeric ||
		g.HorisontalCellsMarker == Alphabetic
	if !checkHorizontalmarker {
		return fmt.Errorf("wrong type of horizontal marker")
	}
	return nil
}

func GenerateGrid(params *GridGeneratorInput) (image.Image, error) {
	if err := params.Valide(); err != nil {
		return nil, err
	}
	maxX := params.CellWidth * params.HorisontalCellsNum
	maxY := params.CellHeight * params.VerticalCellsNum
	if params.FrameWidth > 0 {
		maxX += params.FrameWidth * (params.HorisontalCellsNum + 1)
		maxY += params.FrameWidth * (params.VerticalCellsNum + 1)
	}
	grid := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
	draw.Draw(grid, grid.Bounds(), image.NewUniform(params.BackgroundColor), image.Point{}, draw.Src)

	if params.FrameWidth > 0 {
		for i := 0; i <= params.VerticalCellsNum; i++ {
			DrawHorisontalLine(
				grid,
				0,
				maxX,
				(params.CellHeight+params.FrameWidth)*i,
				params.FrameWidth,
				params.FrameColor,
			)
		}
		for i := 0; i <= params.HorisontalCellsNum; i++ {
			DrawVerticalLine(
				grid,
				(params.CellWidth+params.FrameWidth)*i,
				0,
				maxY,
				params.FrameWidth,
				params.FrameColor,
			)
		}
	}

	if params.MarkerOX == nil && params.MarkerOY == nil {
		return grid, nil
	}

	c := freetype.NewContext()
	c.SetDPI(params.DPI)
	c.SetFont(params.Font)
	c.SetFontSize(params.FontSize)
	c.SetClip(grid.Bounds())
	c.SetDst(grid)
	c.SetSrc(image.Black)
	c.SetHinting(font.HintingNone)

	for i := 0; i < params.HorisontalCellsNum; i++ {
		for j := 0; j < params.VerticalCellsNum; j++ {
			horizontalOffset := (params.FrameWidth+params.CellWidth)*i +
				params.FrameWidth + params.TextPaddingLeft
			verticalOffset := (params.FrameWidth+params.CellHeight)*j +
				params.FrameWidth + params.TextPaddingTop + int(c.PointToFixed(params.FontSize)>>6)
			point := freetype.Pt(horizontalOffset, verticalOffset)
			str := params.MarkerOX.Value(i) + params.Delimeter + params.MarkerOY.Value(j)
			_, err := c.DrawString(str, point)
			if err != nil {
				return nil, err
			}
		}
	}
	return grid, nil
}
