package display

import (
	"log"

	"github.com/nsf/termbox-go"
)

// Display represents a CHIP-8 display
type Display [width][height]bool

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
func (d *Display) DrawSprite(x, y int, s Sprite) bool {
	defer d.refresh()

	collision := false

	for i, v := range s {
		for i2, v2 := range getBitsFromByte(v) {
			x2, y2 := x+i2, i+y

			if x2 > width-1 {
				x2 = x2 - width
			}

			if y2 > height-1 {
				y2 = y2 - height
			}

			current := d[x2][y2]
			new := (v2 == 1)
			d[x2][y2] = current != new

			if !collision && current && !new {
				collision = true
			}
		}
	}

	return collision
}

func (d *Display) refresh() {
	defer termbox.Flush()

	for x := range d {
		for y := range d[x] {
			v := d[x][y]

			if v {
				termbox.SetCell(x, y, ' ', termbox.ColorWhite, termbox.ColorWhite)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorWhite, termbox.ColorDefault)
			}
		}
	}
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
