package main

import "github.com/scottrangerio/go-chip8/display"
import "time"

var zeroSprite = display.Sprite{
	0xF0,
	0x90,
	0x90,
	0x90,
	0xF0,
}

func main() {
	d := new(display.Display)
	d.Init()
	defer d.Close()

	d.DrawSprite(63, 0, zeroSprite)

	time.Sleep(2 * time.Second)
}
