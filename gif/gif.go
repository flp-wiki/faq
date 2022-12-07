package gif

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

func Generate() error {
	dc := gg.NewContext(1200, 628)

	drawBackground(dc)

	if err := drawTitle(dc, "flp.wiki FAQ"); err != nil {
		return errors.Wrap(err, "draw title")
	}

	askI := gg.NewContextForImage(dc.Image())

	if err := ask(askI, "Can I clone my Credit Card with the flipper?"); err != nil {
		return errors.Wrap(err, "draw ask")
	}

	answerI := gg.NewContextForImage(dc.Image())
	if err := answer(answerI, "No, It has encryption why would you want to clone your card? Have you thought about fraud?"); err != nil {
		return errors.Wrap(err, "draw answerer")
	}

	palettedImage1 := image.NewPaletted(askI.Image().Bounds(), palette.Plan9)
	draw.FloydSteinberg.Draw(palettedImage1, askI.Image().Bounds(), askI.Image(), image.ZP)
	palettedImage2 := image.NewPaletted(answerI.Image().Bounds(), palette.Plan9)
	draw.FloydSteinberg.Draw(palettedImage2, answerI.Image().Bounds(), answerI.Image(), image.ZP)
	path := filepath.Join("dist", "image.gif")
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "create gif file")
	}
	gif.EncodeAll(f, &gif.GIF{
		Image: []*image.Paletted{
			palettedImage1,
			palettedImage2,
		},
		Delay: []int{300, 400},
	})

	return nil
}

func drawQuestion(dc *gg.Context, s string) error {

	x := float64(200)
	y := float64(100)
	w := float64(dc.Width()) - 400
	h := float64(dc.Height()) - 300

	dc.DrawRoundedRectangle(x, y, w, h, 20)
	dc.SetLineWidth(4)
	dc.Stroke()
	textColor := color.Black
	fontPath := filepath.Join("assets", "haxrcorp_4089_cyrillic_altgr.ttf")
	if err := dc.LoadFontFace(fontPath, 90); err != nil {
		return errors.Wrap(err, "load haxrcorp_4089_cyrillic_altgr")
	}

	dc.SetColor(textColor)
	dc.DrawStringWrapped(s, x, (h / 2), 0, 0, w, 1.2, gg.AlignCenter)
	return nil
}

func watermark(dc *gg.Context) error {
	textColor := color.Black
	fontPath := filepath.Join("assets", "helvb08.ttf")
	if err := dc.LoadFontFace(fontPath, 60); err != nil {
		return errors.Wrap(err, "load helvb08")
	}
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	dc.SetColor(mutedColor)
	marginY := float64(30)
	s := "https://flp.wiki/"
	_, textHeight := dc.MeasureString(s)
	x := float64(70)
	y := float64(dc.Height()) - textHeight - marginY
	dc.DrawString(s, x, y)
	return nil
}

func drawBackground(dc *gg.Context) {
	dc.SetColor(color.RGBA{0xFF, 0x82, 0x00, 0xFF})
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.Fill()
}

func drawTitle(dc *gg.Context, s string) error {

	x, y := float64(0), float64(0)
	fontPath := filepath.Join("assets", "Born2bSportyV2.ttf")
	if err := dc.LoadFontFace(fontPath, 80); err != nil {
		return errors.Wrap(err, "load font")
	}
	dc.SetColor(color.Black)
	marginX := 20.0
	marginY := 10.0
	textWidth, textHeight := dc.MeasureString(s)
	x = float64(dc.Width()) - textWidth - marginX
	y = textHeight + marginY
	dc.DrawString(s, x, y)
	return nil
}

func ask(dc *gg.Context, question string) error {
	fn := filepath.Join("assets", "ask.png")
	i, err := gg.LoadImage(fn)
	si := imaging.Resize(i, i.Bounds().Dx()*4, i.Bounds().Dy()*4, imaging.NearestNeighbor)
	if err != nil {
		return errors.Wrap(err, "load image")
	}
	x := 0
	y := dc.Height() - si.Bounds().Dy()
	dc.DrawImage(si, x, y)

	if err := drawQuestion(dc, question); err != nil {
		return errors.Wrap(err, "draw question")
	}

	return nil
}

func answer(dc *gg.Context, answer string) error {
	fn := filepath.Join("assets", "answer.png")
	i, err := gg.LoadImage(fn)
	si := imaging.Resize(i, i.Bounds().Dx()*4, i.Bounds().Dy()*4, imaging.NearestNeighbor)
	if err != nil {
		return errors.Wrap(err, "load image")
	}
	x := dc.Width() - si.Bounds().Dx()
	y := dc.Height() - si.Bounds().Dy()
	dc.DrawImage(si, x, y)

	if err := drawQuestion(dc, answer); err != nil {
		return errors.Wrap(err, "draw question")
	}

	return nil
}
