package main

func main() {
	bs := make(bubbles)
	bs["1"] = newBubble("hello", 16)
	bs["2"] = newBubble("world", 16)
	bs["3"] = newBubble("*glomps u*", 16)
	bs["4"] = newBubble("penis", 16)
	bs["5"] = newBubble("lol", 16)
	bs["6"] = newBubble("colada", 16)
	bs["7"] = newBubble("oh, woah", 16)
	bs["8"] = newBubble("what's this?", 16)
	bs["9"] = newBubble("best girl", 16)
	bs["10"] = newBubble("icetal", 16)
	bs["11"] = newBubble(":jadeteefs:", 16)
	bs["12"] = newBubble("oct2pus", 16)

	bs.draw(500, 500).SavePNG("hello.png")

}
