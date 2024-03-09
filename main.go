package main

func main() {
	bub := newBubble("hello world!", 16)
	bub.DC.SavePNG("hello.png")
}
