package main

import (
	"fmt"
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
	op := cpu.Run()

	fmt.Printf("%04X\n", op)
}
