package shapes

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"srcgo/github.com/antoniofmoliveira/patterns/behavioral/strategy/strategy"
)

type ImageSquare struct {
	strategy.DrawOutput
}

func (t *ImageSquare) Draw() error {
	width := 800
	height := 600
	origin := image.Point{0, 0}
	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})
	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)
	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)

	if t.Writer == nil {
		return fmt.Errorf("no writer stored on ImageSquare")
	}
	if err := jpeg.Encode(t.Writer, bgImage, quality); err != nil {
		return fmt.Errorf("error writing image to disk")
	}
	if t.LogWriter != nil {
		io.Copy(t.LogWriter, bytes.NewReader([]byte("Image written in provided writer\n")))
	}
	return nil
}
