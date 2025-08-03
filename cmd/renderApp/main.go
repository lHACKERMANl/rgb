package main

import (
	"github.com/lHACKERMANl/rgb/internal/app"
	"log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	app := app.NewApp()
	if err := app.Run(); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}
