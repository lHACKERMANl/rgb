package main

import "github.com/lHACKERMANl/rgb/internal/shaderService/cmd/shaderService/app"

func mian() {
	err := app.NewApp().Run()
	if err != nil {
		panic(err)
	}
}
