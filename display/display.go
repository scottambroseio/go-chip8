package display

import "github.com/nsf/termbox-go"
import "log"

type Display [64][32]bool
type Sprite []byte

// Init initializing the display
func (d *Display) Init() {
	err := termbox.Init()

	if err != nil {
		log.Fatal(err)
	}
}

// DrawSprite draws a sprite to the display
func (d *Display) DrawSprite(s Sprite) {
	for i, v := range s {
		for i2, v2 := range getBitsFromByte(v) {
			if v2 == 1 {
				termbox.SetCell(i2, i, ' ', termbox.ColorWhite, termbox.ColorWhite)
			} else {
				termbox.SetCell(i2, i, ' ', termbox.ColorWhite, termbox.ColorDefault)
			}
		}
	}

	termbox.Flush()
}

// Close terminates and shuts down the display
func (d Display) Close() {
	termbox.Close()
}

func getBitsFromByte(b byte) [8]byte {
	return [8]byte{
		(b << 0) >> 7,
		(b << 1) >> 7,
		(b << 2) >> 7,
		(b << 3) >> 7,
		(b << 4) >> 7,
		(b << 5) >> 7,
		(b << 6) >> 7,
		(b << 7) >> 7,
	}
}
