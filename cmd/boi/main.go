package main

import (
	"fmt"
	"os"

	"github.com/Laellekoenig/boi/internal/emulator"
	"github.com/Laellekoenig/boi/internal/webserver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: boi rom.go")
		return
	}

	filename := os.Args[1]
	e := emulator.NewEmulator(filename)
	s := webserver.NewServer(e)
	s.Logger.Fatal(s.Start(":3000"))
}
