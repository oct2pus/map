package main

import (
	"github.com/fogleman/gg"
)

func main() {
	dc := drawBubble("Hello World!", 16)
	dc.SavePNG("hello.png")
}

// drawBubble draws text with a bubble around it.
// content must be greater than 0.
func drawBubble(s string, pt float64) *gg.Context {
	inch := 72.272                      //1 inch is 72.272 (font) points. or 72 points. im getting conflicting information here.
	px := inch / 96                     //1 px is 1/96th of an inch. According to some random website.
	padX, padY := 0.0, 0.0              // left in case I change my mind on absolute padding vs relative padding
	size := float64(len(s)) * (px * pt) /* this was incorrect but it does make for a nice amount of padding per text.
	It gets wider as you add more text. This is because the program should be cursed. */
	x, y := size+padX, (pt*2)+padY

	dc := gg.NewContext(int(x), int(y))
	err := dc.LoadFontFace("./Hack-Regular.ttf", pt)
	if err != nil {
		print("lol")
	}
	dc.DrawRoundedRectangle(0, 0, x, y, 10)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	dc.DrawRoundedRectangle(0, 0, x, y, 10)
	dc.SetRGB(256, 256, 256)
	dc.SetLineWidth(3)
	dc.Stroke()
	dc.DrawStringAnchored(s, x/2, y/2, 0.5, 0.5)
	return dc
}
