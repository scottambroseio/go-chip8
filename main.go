package main

import (
	"io/ioutil"
	"log"

	"github.com/scottrangerio/go-chip8/cpu"
)

func main() {
	cpu := cpu.NewCPU()

	rom, err := ioutil.ReadFile("./roms/pong.rom")

	if err != nil {
		log.Fatal(err)
	}

	cpu.LoadRom(rom)
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	termbox.Close()
	// 	os.Exit(0)
	// }()
	cpu.Run()
}
