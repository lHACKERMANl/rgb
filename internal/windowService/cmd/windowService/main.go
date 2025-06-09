package main

import "github.com/lHACKERMANl/rgb/windowService/cmd/windowService/app"

func mian() {
	err := app.NewApp().Run()
	if err != nil {
		panic(err)
	}
}
