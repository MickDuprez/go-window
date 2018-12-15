package main

import (
	"fmt"

	"golang.org/x/mobile/event/paint"

	"golang.org/x/mobile/event/mouse"

	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/size"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

// set up some Global var's
var winWidth, winHeight = 800, 650

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title:  "Simple Window for Graphics",
			Width:  winWidth,
			Height: winHeight,
		})
		if err != nil {
			fmt.Printf("Failed to create Window - %v", err)
			return
		}
		defer w.Release()

		var cnt int // counter to help with messages
		for {
			// Handle some main window events to do any saving of resources
			// and final clean up etc as required.
			switch e := w.NextEvent().(type) {

			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					// Do any final cleanup or saving here:
					// ---

					fmt.Println("Thanks for playing, goodbye!")
					return // quit the application.
				}

			// let's handle the window resize event, we'll need this to update
			// the global winWidth/winHeight var's for use in building buffers
			// to suit the window.
			case size.Event:
				cnt++

				fmt.Printf("Event %3d: size.Event Size = %v\n", cnt, e.Size())
				winWidth = e.Size().X
				winHeight = e.Size().Y

			// handle keyboard events:
			case key.Event:
				cnt++
				// print out the current key being pressed, notice it sends 2 messages
				// one for press, one for release. For holding a key press it sends 'none'
				// uncomment the below to test and see.
				fmt.Printf("Event %3d: key.Event - Key %v event was %v\n", cnt, e.Code, e.Direction)

				// lets handle some specific scenarios:
				switch e.Code {
				case key.CodeEscape:
					// let's use this to quit for now:
					fmt.Println("\n\nYou have escaped the window, bye!")
					return

				case key.CodeA:
					if e.Direction == key.DirPress {
						fmt.Println("  moving one step left.")
					}
					if e.Direction == key.DirNone {
						fmt.Println("    he's still going!")
					}
					if e.Direction == key.DirRelease {
						fmt.Println(" Phew!, stopped going left.")
					}
				} // --- end key switch

			// handle mouse events, can we see a pattern here? ;)
			case mouse.Event:
				switch e.Button {
				case mouse.ButtonLeft:
					if e.Direction == mouse.DirPress {
						fmt.Println("left mouse button down")
					}
					// this doesn't fire for some reason??
					if e.Direction == mouse.DirNone {
						fmt.Println("left mouse button being held")
					}
					if e.Direction == mouse.DirRelease {
						fmt.Println("left mouse button released")
					}

				case mouse.ButtonRight:
					// handle press as single step (not working in Windows??):
					if e.Direction == mouse.DirStep {
						fmt.Println("right mouse button pressed and released in single step.")
					}
					if e.Direction == mouse.DirPress {
						fmt.Println("right mouse button pressed.")
					}
				case mouse.ButtonMiddle:
					if e.Direction == mouse.DirPress {
						fmt.Println("middle mouse button pressed.")
					}
				// doesn't work if you have set mouse wheel to middle button.
				case mouse.ButtonWheelDown:
					if e.Direction == mouse.DirPress {
						fmt.Println("mouse wheel button pressed.")
					}
				case mouse.ButtonWheelUp:
					fmt.Println("mouse wheel direction UP.")
					// let's try calling an event manually, the Paint event
					// could well be useful in future:
					w.Send(paint.Event{true})
				} // --- end mouse switch

			case paint.Event:
				// may not need/be able to use this, see below for details
				// https://godoc.org/golang.org/x/mobile/event/paint#Event
				fmt.Printf("paint.Event called: %v\n", e.External)

			} // --- end events
		} // --- end app loop
	})
}
