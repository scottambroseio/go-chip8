package cpu

import (
	"math/rand"
	"time"

	"github.com/scottrangerio/go-chip8/cpu/opcode"
	"github.com/scottrangerio/go-chip8/display"
	"github.com/scottrangerio/go-chip8/sprites"
)

// LoadRom loads a rom into memory
func (c *CPU) LoadRom(d []byte) {
	s := 0x200

	for i := 0; i < len(d); i++ {
		c.memory[s+i] = d[i]
	}
}

func (c *CPU) getOpcode() opcode.Opcode {
	return opcode.NewOpcode(c.memory[c.pc], c.memory[c.pc+1])
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
	st     byte
	dt     byte
}

// NewCPU creates and initializes a new CPU
func NewCPU() *CPU {
	cpu := &CPU{
		pc: 0x200,
	}

	for i, c := 0, 0; i < 16; i++ {
		for j := 0; j < 5; j++ {
			cpu.memory[c+i+j] = sprites.Sprites[i][j]
		}
		c += 4
	}

	return cpu
}

// Run runs the emulator
func (c *CPU) Run() opcode.Opcode {
	d := new(display.Display)
	d.Init()
	defer d.Close()

	for {
		op := c.getOpcode()

		if c.dt > 0 {
			c.dt--
		}

		if c.st > 0 {
			c.st--
		}

		switch op.LeadByte() {
		case 0x0:
			c.pc = c.stack[c.sp]
			c.sp--
			c.pc += 2
		case 0x1:
			c.pc = op.NNN()
		case 0x2:
			c.sp++
			c.stack[c.sp] = c.pc
			c.pc = op.NNN()
		case 0x3:
			x := op.X()
			if c.v[x] == op.KK() {
				c.pc += 2
			}
			c.pc += 2
		case 0x4:
			x := op.X()
			if c.v[x] != op.KK() {
				c.pc += 2
			}
			c.pc += 2
		case 0x6:
			x := op.X()
			c.v[x] = op.KK()
			c.pc += 2
		case 0x7:
			x := op.X()
			c.v[x] = c.v[x] + op.KK()
			c.pc += 2
		case 0x8:
			switch op.N() {
			case 0x0:
				x := op.X()
				y := op.Y()
				c.v[x] = c.v[y]
				c.pc += 2
			case 0x1:
				x := op.X()
				y := op.Y()
				r := c.v[x] | c.v[y]
				c.v[x] = r
				c.pc += 2
			case 0x2:
				x := op.X()
				y := op.Y()
				r := c.v[x] & c.v[y]
				c.v[x] = r
				c.pc += 2
			case 0x3:
				x := op.X()
				y := op.Y()
				r := c.v[x] ^ c.v[y]
				c.v[x] = r
				c.pc += 2
			case 0x4:
				x := op.X()
				y := op.Y()
				r := uint16(c.v[x]) + uint16(c.v[y])
				if r > 0xFF {
					c.v[0xF] = 1
				} else {
					c.v[0xF] = 0
				}
				c.v[x] = byte(r | 0x00)
				c.pc += 2
			case 0x5:
				x := op.X()
				y := op.Y()
				if c.v[x] > c.v[y] {
					c.v[0xF] = 1
				} else {
					c.v[0xF] = 0
				}
				c.v[x] = c.v[x] - c.v[y]
				c.pc += 2
			case 0x6:
				x := op.X()
				c.v[0xF] = c.v[x] & 0x01
				c.v[x] /= 2
				c.pc += 2
			default:
				time.Sleep(1 * time.Second)
				return op
			}
		case 0xA:
			c.i = op.NNN()
			c.pc += 2
		case 0xC:
			x := op.X()
			rand.Seed(time.Now().Unix())
			r := rand.Intn(255)
			c.v[x] = byte(r) & op.KK()
			c.pc += 2
		case 0xD:
			x := c.v[op.X()]
			y := c.v[op.Y()]
			n := uint16(op.N())

			b := c.memory[c.i : c.i+n]

			d.DrawSprite(int(x), int(y), b)
			time.Sleep((1000 / 60) * time.Millisecond)
			c.pc += 2
		case 0xE:
			switch op.KK() {
			case 0x00A1:
				rand.Seed(time.Now().Unix())

				if rand.Intn(2) == 1 {
					c.pc += 2
				}
				c.pc += 2
			case 0x009E:
				rand.Seed(time.Now().Unix())
				if rand.Intn(2) == 1 {
					c.pc += 2
				}
				c.pc += 2
			}
		case 0xF:
			switch op.KK() {
			case 0x0007:
				x := op.X()
				c.v[x] = c.dt
				c.pc += 2
			case 0x0015:
				x := op.X()
				c.dt = c.v[x]
				c.pc += 2
			case 0x0018:
				x := op.X()
				c.st = c.v[x]
				c.pc += 2
			case 0x0029:
				x := op.X()
				v := c.v[x] * 0x05
				c.i = uint16(v)
				c.pc += 2
			case 0x0033:
				x := op.X()
				c.memory[c.i] = c.v[x] / 100
				c.memory[c.i+1] = (c.v[x] / 10) % 10
				c.memory[c.i+2] = (c.v[x] % 100) % 10

				c.pc += 2
			case 0x0065:
				x := uint16(op.X())

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
