package ui

import (
	"bytes"
	"image/color"
	"log"
	"quizapp/colors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	uiImage           *ebiten.Image
	FaceSourceBold    *text.GoTextFaceSource
	FaceSourceRegular *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(gobold.TTF))
	if err != nil {
		log.Fatal(err)
	}
	FaceSourceBold = s

	s, err = text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	FaceSourceRegular = s
}

// Button represents customizable UI element that triggers specified action.
type Button struct {
	Background *ebiten.Image
	Color      color.RGBA
	Text       string
	Width      int
	Height     int
	Coord

	Hover   bool
	BorderW int
	Action  func()

	fs       *text.GoTextFaceSource
	fontSize float64
}

// Draw renders button on the destination image.
func (b *Button) Draw(dst *ebiten.Image) {
	b.Background.Fill(b.Color)

	opText := &text.DrawOptions{}
	opText.GeoM.Translate(float64(b.Width)/2, float64(b.Height)/2)
	opText.ColorScale.ScaleWithColor(color.White)
	opText.PrimaryAlign = text.AlignCenter
	opText.SecondaryAlign = text.AlignCenter
	opText.LineSpacing = 26
	text.Draw(b.Background, b.Text, &text.GoTextFace{
		Source: b.fs,
		Size:   b.fontSize,
	}, opText)

	opImg := &ebiten.DrawImageOptions{}
	opImg.GeoM.Translate(float64(b.X), float64(b.Y))

	if b.Hover {
		vector.StrokeRect(b.Background, 1, 1, float32(b.Width)-1, float32(b.Height)-1, 3, colors.Orange, false)
	}

	dst.DrawImage(b.Background, opImg)

}

// Update updates button's state based on the current cursor position.
func (b *Button) Update() {
	x, y := ebiten.CursorPosition()

	if x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height {
		b.Hover = true
	} else {
		b.Hover = false
	}

	b.click()
}

// click checks whether mouse cursor is over the button and handles the button click event
func (b *Button) click() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(ebiten.MouseButtonLeft)) && b.Hover {
		b.Action()
	}
}

func NewButton(w, h int, c color.RGBA, t string, fs *text.GoTextFaceSource, s float64, posX, posY int, a func()) *Button {
	b := Button{
		Color:      c,
		Text:       t,
		Width:      w,
		Height:     h,
		Background: ebiten.NewImage(w, h),
		Coord:      Coord{posX, posY},
		Action:     a,
		fs:         fs,
		fontSize:   s,
	}

	return &b
}
