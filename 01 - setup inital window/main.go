package main

import (
	"fmt"

	"golang.org/x/mobile/event/lifecycle"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

func main() {
	// the driver package calls its Main to provide access to the screen
	// through the OS drivers abstracting away the underlying system.
	driver.Main(func(s screen.Screen) {
		// using the screen, create a top level double buffered window:
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title:  "Simple Window for Graphics",
			Width:  800,
			Height: 650,
		})
		if err != nil {
			fmt.Printf("Failed to create Window - %v", err)
			return
		}
		defer w.Release()

		// We have a window, now we need a loop to handle window events
		// in regards to other windows in the OS
		var cnt int // counter to help with messages
		for {
			switch e := w.NextEvent().(type) {

			case lifecycle.Event:
				cnt++
				fmt.Printf("Event %d: From %s To %s\n", cnt, e.From, e.To)
				if e.To == lifecycle.StageDead {
					fmt.Println("lifecycle is StageDead, goodbye!")
					return // quit the application.
				}

				// the following 2 Stages seem to be the only ones that work
				// as you would expect. StageInvisible doesn't work on Windows OS
				// but this may be different on other systems.
				if e.To == lifecycle.StageFocused {
					fmt.Println("window now has the focus")
				}
				if e.From == lifecycle.StageFocused {
					fmt.Println("window has lost the focus")
				}
				// we'll add other events as they are discovered...
			}
		}

	})
}
