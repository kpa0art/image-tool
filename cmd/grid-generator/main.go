package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/kpa0art/image-tool/pkg/tool"
)

var (
	dpi                   = flag.Float64("dpi", 72, "Разрешение экрана в точках на дюйм.")
	verticalCellsNum      = flag.Int("vcells", 1, "Количество ячеек по вертикали в сетке")
	horisontalCellsNum    = flag.Int("hcells", 1, "Количество ячеек по горизонтали в сетке")
	cellheight            = flag.Int("height", 200, "Высота одной ячейки")
	cellwidth             = flag.Int("width", 200, "Ширина одной ячейки")
	fontSize              = flag.Int("font-size", 12, "use to place a text in the image")
	fontColorHex          = flag.String("font-color-gex", "000000", "Шестнадцатеричное (hex) представление цвета шрифта текста.")
	cellFrameWidth        = flag.Int("frame-width", 1, "Толщина рамки ячейки в пикселях")
	cellFrameColorHex     = flag.String("frame-color-hex", "000000", "Шестнадцатеричное (hex) представление цвета рамки ячейки")
	cellFrameTransparent  = flag.Int("frame-transparent", 255, "Прозрачность рамки ячейки. Возможные значения в диапазоне [0; 255]. Если установить в 0, то рамка ячейки будет полностью прозрачной (невидимой).")
	backgroundColorHex    = flag.String("bg-color-hex", "FFFFFF", "Шестнадцатеричное (hex) представление цвета заднего плана ячейки")
	backgroundTransparent = flag.Int("bg-transparent", 255, "Прозрачность заднего плана ячейки. Возможные значения в диапазоне [0; 255]. Если установить в 0, то задний план будет полностью прозрачным (невидимым).")
	filename              = flag.String("filename", "grid_{datetime}.png", "Имя файла, включая полный путь к нему.")
	textPaddingLeft       = flag.Int("padding-left", 10, "Отступ слева от рамки ячейки для текста.")
	textPaddingTop        = flag.Int("padding-top", 10, "Отступ сверху от рамки ячейки для текста.")
	fontFileName          = flag.String("font-file", "no-file", "Название файла шрифта, включая полный путь к нему. Обычно такие файлы имеют формат '.ttf'")
)

func generateFileName() string {
	return "grid_" + time.Now().Format("20060102150405") + ".png"
}

func getFont(fontFileName string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(fontFileName)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	return font, nil
}

func main() {
	flag.Parse()

	backgroundColor, err := tool.Hex2RGBA(*backgroundColorHex)
	if err != nil {
		log.Fatal(err)
	}
	backgroundColor.A = uint8(*backgroundTransparent)

	frameColor, err := tool.Hex2RGBA(*cellFrameColorHex)
	if err != nil {
		log.Fatal(err)
	}
	frameColor.A = uint8(*cellFrameTransparent)

	font, err := getFont(*fontFileName)
	if err != nil {
		log.Fatal(err)
	}
	fontColor, err := tool.Hex2RGBA(*fontColorHex)
	if err != nil {
		log.Fatal(err)
	}

	image, err := tool.GenerateGrid(&tool.GridGeneratorInput{
		VerticalCellsNum:      *verticalCellsNum,
		HorisontalCellsNum:    *horisontalCellsNum,
		CellHeight:            *cellheight,
		CellWidth:             *cellwidth,
		VerticalCellsMarker:   tool.Numeric,
		HorisontalCellsMarker: tool.Numeric,
		BackgroundColor:       backgroundColor,
		FrameColor:            frameColor,
		FrameWidth:            *cellFrameWidth,
		Font:                  font,
		FontColor:             fontColor,
		FontSize:              float64(*fontSize),
		TextPaddingLeft:       *textPaddingLeft,
		TextPaddingTop:        *textPaddingTop,
		DPI:                   *dpi,
		MarkerOY:              tool.NumericMarker{},
		MarkerOX:              tool.AlphabeticMarker{},
	})
	if err != nil {
		log.Fatal(err)
	}
	if *filename == "grid_{datetime}.png" {
		*filename = generateFileName()
	}
	outFile, err := os.Create(*filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	writer := bufio.NewWriter(outFile)
	err = png.Encode(writer, image)
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %s OK.\n", *filename)
}
