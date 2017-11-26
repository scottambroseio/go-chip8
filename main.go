package main

import (
	"io/ioutil"
	"log"
	"time"

	termbox "github.com/nsf/termbox-go"
	"github.com/scottrangerio/go-chip8/cpu"
)

func main() {
	cpu := cpu.NewCPU()

	rom, err := ioutil.ReadFile("./roms/pong.rom")

	if err != nil {
		log.Fatal(err)
	}

	cpu.LoadRom(rom)

	done := make(chan struct{})
	q := make(chan termbox.Event)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				q <- termbox.PollEvent()
			}
		}
	}()

	kb := make(map[byte]bool)

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				if e := <-q; e.Type == termbox.EventKey {
					if e.Key == termbox.KeyEsc {
						done <- struct{}{}
					}
					switch e.Ch {
					case 'w':
						kb[0x1] = !kb[0x1]
						go func() {
							time.Sleep(150 * time.Millisecond)
							kb[0x1] = !kb[0x1]
						}()
					case 's':
						kb[0x4] = !kb[0x4]
						go func() {
							time.Sleep(150 * time.Millisecond)
							kb[0x4] = !kb[0x4]
						}()
					case 'u':
						kb[0xC] = !kb[0xC]
						go func() {
							time.Sleep(150 * time.Millisecond)
							kb[0xC] = !kb[0xC]
						}()
					case 'j':
						kb[0xD] = !kb[0xD]
						go func() {
							time.Sleep(150 * time.Millisecond)
							kb[0xD] = !kb[0xD]
						}()
						// case '1':
						// 	kb[0x1] = !kb[0x1]
						// 	go func() {
						// 		time.Sleep(150 * time.Millisecond)
						// 		kb[0x1] = !kb[0x1]
						// 	}()
						// case '2':
						// case '3':
						// case '4':
						// 	kb[0xC] = !kb[0xC]
						// 	go func() {
						// 		time.Sleep(150 * time.Millisecond)
						// 		kb[0xC] = !kb[0xC]
						// 	}()
						// case 'q':
						// 	kb[0x4] = !kb[0x4]
						// 	go func() {
						// 		time.Sleep(150 * time.Millisecond)
						// 		kb[0x4] = !kb[0x4]
						// 	}()
						// case 'w':
						// case 'e':
						// case 'r':
						// 	kb[0xD] = !kb[0xD]
						// 	go func() {
						// 		time.Sleep(150 * time.Millisecond)
						// 		kb[0xD] = !kb[0xD]
						// 	}()
						// case 'a':
						// case 's':
						// case 'd':
						// case 'f':
						// case 'z':
						// case 'x':
						// case 'c':
						// case 'v':
					}
				}
			}

		}
	}()

	cpu.Run(done, kb)
}
