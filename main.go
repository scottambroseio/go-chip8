package main

import (
	"io/ioutil"
	"log"
	"os"
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

	// q := make(chan termbox.Event)
	// go func() {
	// 	for {
	// 		q <- termbox.PollEvent()
	// 	}
	// }()
	// termbox.Init()
	// var log []interface{}
	// for i := 0; i < 5; i++ {
	// 	select {
	// 	case e := <-q:
	// 		log = append(log, e.Type)
	// 		break
	// 	default:
	// 		log = append(log, "no event")
	// 	}
	// 	time.Sleep(1 * time.Second)
	// }
	// termbox.Close()

	// for _, v := range log {
	// 	fmt.Println(v)
	// }
	go func() {
		time.Sleep(10 * time.Second)
		termbox.Close()
		os.Exit(0)
	}()
	cpu.Run()
}
