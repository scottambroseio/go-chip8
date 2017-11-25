package display

import "github.com/nsf/termbox-go"
import "log"

// Display represents a CHIP-8 display
type Display [height][width]bool

// Sprite represents a CHIP-8 sprite
type Sprite []byte

const width = 64
const height = 32

// Init initializing the display
func (d *Display) Init() {
	err := termbox.Init()

	if err != nil {
		log.Fatal(err)
	}
}

// DrawSprite draws a sprite to the display
func (d *Display) DrawSprite(x, y int, s Sprite) {
	for i, v := range s {
		for i2, v2 := range getBitsFromByte(v) {
			x2, y2 := x+i2, i+y

			if x2 > width-1 {
				x2 = x2 - width
			}

			if y2 > height-1 {
				y2 = y2 - height
			}

			if v2 == 1 {
				termbox.SetCell(x2, y2, ' ', termbox.ColorWhite, termbox.ColorWhite)
			} else {
				termbox.SetCell(x2, y2, ' ', termbox.ColorWhite, termbox.ColorDefault)
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
