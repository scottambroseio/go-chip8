package cpu

import (
	"time"

	"github.com/scottrangerio/go-chip8/sprites"

	"github.com/scottrangerio/go-chip8/display"
)

// LoadRom loads a rom into memory
func (c *CPU) LoadRom(d []byte) {
	s := 0x200

	for i := 0; i < len(d); i++ {
		c.memory[s+i] = d[i]
	}
}

func (c *CPU) getOpcode() uint16 {
	return uint16(c.memory[c.pc])<<8 | uint16(c.memory[c.pc+1])
}

// CPU represents a CHIP-8 CPU
type CPU struct {
	v      [16]byte
	sp     byte
	pc     uint16
	i      uint16
	stack  [16]uint16
	sound  byte
	timer  byte
	memory [4096]byte
}

// NewCPU creates and initializes a new CPU
func NewCPU() *CPU {
	cpu := &CPU{
		pc: 0x200,
	}

	for i := 0; i < 16; i += 5 {
		for j := 0; j < 5; j++ {
			cpu.memory[i+j] = sprites.Sprites[i][j]
		}
	}

	return cpu
}

// Run runs the emulator
func (c *CPU) Run() uint16 {
	d := new(display.Display)
	d.Init()
	defer d.Close()

	for {
		op := c.getOpcode()

		switch op & 0xF000 {
		case 0x0000:
			c.pc = c.stack[c.sp]
			c.sp--
			c.pc += 2
		case 0x2000:
			c.sp++
			c.stack[c.sp] = c.pc
			c.pc = op & 0x0FFF
		case 0x6000:
			x := (op & 0x0F00) >> 8
			kk := byte(op & 0x00FF)
			c.v[x] = kk
			c.pc += 2
		case 0x7000:
			x := (op & 0x0F00) >> 8
			kk := byte(op & 0x00FF)
			c.v[x] = c.v[x] + kk
			c.pc += 2
		case 0xA000:
			c.i = op & 0x0FFF
			c.pc += 2
		case 0xD000:
			x := c.v[(op&0x0F00)>>8]
			y := c.v[(op&0x00F0)>>4]
			n := op & 0x000F

			b := c.memory[c.i : c.i+n]

			d.DrawSprite(int(x), int(y), b)
			c.pc += 2
		case 0xF000:
			switch op & 0x00FF {
			case 0x0029:
				x := (op & 0x0F00) >> 8
				v := c.v[x]
				c.i = uint16(v)
				c.pc += 2
			case 0x0033:
				x := (op & 0x0F00) >> 8
				c.memory[c.i] = c.v[x] / 100
				c.memory[c.i+1] = (c.v[x] / 10) % 10
				c.memory[c.i+2] = (c.v[x] % 100) % 10

				c.pc += 2
			case 0x0065:
				x := (op & 0x0F00) >> 8

				for i := uint16(0); i <= x; i++ {
					c.v[i] = c.memory[c.i+i]
				}
				c.pc += 2
			default:
				time.Sleep(1 * time.Second)
				return op
			}
		default:
			time.Sleep(1 * time.Second)
			return op
		}
	}
}