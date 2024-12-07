package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// TextBox represents graphical text box with background.
type TextBox struct {
	Background *ebiten.Image
	Color      color.RGBA
	Text       string
	Width      int
	Height     int
	Coord
}

// Draw renders the text box on the destination image.
func (tb *TextBox) Draw(dst *ebiten.Image) {
	tb.Background.Fill(tb.Color)
	// TODO: text wrapping
	opText := &text.DrawOptions{}
	opText.ColorScale.ScaleWithColor(color.White)
	opText.GeoM.Translate(float64(tb.Width)/2, float64(tb.Height)/2)
	opText.ColorScale.ScaleWithColor(color.White)
	opText.PrimaryAlign = text.AlignCenter
	opText.SecondaryAlign = text.AlignCenter
	opText.LineSpacing = 34
	text.Draw(tb.Background, tb.Text, &text.GoTextFace{
		Source: FaceSourceRegular,
		Size:   28,
	}, opText)

	opImg := &ebiten.DrawImageOptions{}
	opImg.GeoM.Translate(float64(tb.X), float64(tb.Y))
	dst.DrawImage(tb.Background, opImg)
}

func NewTextBox(w, h int, c color.RGBA, t string, posX, posY int) *TextBox {
	tb := TextBox{
		Background: ebiten.NewImage(w, h),
		Color:      c,
		Text:       t,
		Width:      w,
		Height:     h,
		Coord:      Coord{posX, posY},
	}

	return &tb
}
