package main

import "github.com/fogleman/gg"

// Grid is the represents the base image that houses various Bubbles and Lines.
type Grid struct {
	DC      *gg.Context
	X, Y    int
	Bubbles [](*bubble)

	Coord [][]bool
}

func newGrid(x, y int) Grid {
	var g Grid
	g.X = x
	g.Y = y
	g.Bubbles = make([]*bubble, 0)

	g.Coord = make([][]bool, y) //the initial value of a bool is false
	for i := range x {
		g.Coord[i] = make([]bool, x)
	}

	return g
}

type bubble struct {
	DC              *gg.Context
	SizeX           int
	SizeY           int
	CoordX, CoordX2 int
	CoordY, CoordY2 int
}

// newBubble creates a bubble object.
// content must be greater than 0.
func newBubble(s string, pt float64) *bubble {
	var b bubble
	inch := 72.272                      //1 inch is 72.272 (font) points. or 72 points. im getting conflicting information here.
	px := inch / 96                     //1 px is 1/96th of an inch. According to some random website.
	padX, padY := 0.0, 0.0              // left in case I change my mind on absolute padding vs relative padding
	to := 25                            // round up to the nearest 25th.
	size := float64(len(s)) * (px * pt) /* this was incorrect but it does make for a nice amount of padding per text.
	It gets wider as you add more text. This is because the program should be cursed. */
	bx, by := size+padX, (pt*2)+padY         // box x and box y
	x, y := roundTo(bx, to), roundTo(by, to) //dc x and dc y

	b.DC = drawBubble(s, pt, x, y, bx, by)
	b.SizeX = int(x)
	b.SizeY = int(y)
	b.CoordX = 0
	b.CoordY = 0
	b.CoordX2 = b.SizeX
	b.CoordY2 = b.SizeY

	return &b
}

// drawBubble creates a gg.Context for a bubble. the size is moved to the nearest 25th pixel and the bubble rests in the center.
func drawBubble(s string, pt, x, y, bx, by float64) *gg.Context {
	dc := gg.NewContext(int(x), int(y))
	err := dc.LoadFontFace("./Hack-Regular.ttf", pt)
	if err != nil {
		print("lol")
	}

	dc.DrawRoundedRectangle((x-bx)/2, (y-by)/2, bx, by, 10)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	dc.DrawRoundedRectangle((x-bx)/2, (y-by)/2, bx, by, 10)
	dc.SetRGB(256, 256, 256)
	dc.SetLineWidth(3)
	dc.Stroke()
	dc.DrawStringAnchored(s, x/2, y/2, 0.5, 0.5)
	return dc
}

// Move moves a bubble to a new coordinate value.
func (b *bubble) Move(x, y int) {
	b.CoordX, b.CoordX2 = b.CoordX+x, b.CoordX2+x
	b.CoordY, b.CoordY2 = b.CoordY+y, b.CoordY2+y
}

// roundTo rounds base up to the next multiple of To.
func roundTo(base float64, To int) float64 {
	i := int(base)
	for i%To != 0 {
		i++
	}
	return float64(i)
}
