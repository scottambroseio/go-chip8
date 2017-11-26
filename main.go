package main

import (
	"io/ioutil"
	"log"

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

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				switch e := <-q; e.Type {
				case termbox.EventKey:
					if e.Key == termbox.KeyEsc {
						done <- struct{}{}
					}
				}
			}

		}
	}()

	cpu.Run(done)
}
