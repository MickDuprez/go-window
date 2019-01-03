package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"

	"golang.org/x/mobile/event/mouse"

	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/size"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

var (
	// set up some global helper var's
	winWidth, winHeight = 800, 650

	// We can get info from the event.Size() function along with other
	// helpful functions and data.
	sizeEvent size.Event
)

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

		// As the size.Event might not be called at startup
		// set them the screenSize the same as our winHeight/winWidth
		screenSize := image.Point{}
		if sizeEvent.Bounds().Max.X == 0 {
			screenSize = image.Point{winWidth, winHeight}
		} else {
			screenSize = image.Point{sizeEvent.WidthPx, sizeEvent.HeightPx}
		}
		screenBuffer, err := s.NewBuffer(screenSize)
		if err != nil {
			log.Fatalf("%v - failed to create screen buffer", err)
		}
		defer screenBuffer.Release()
		// pixBuffer is like a 'back bufffer' and it's what we do our drawing to.
		pixBuffer := screenBuffer.RGBA()

		for {
			// Handle window events:
			switch e := w.NextEvent().(type) {

			case size.Event:
				sizeEvent = e
				// we need to create a new screen buffer, there's no way to resize the old one
				screenBuffer.Release()
				screenBuffer, err = s.NewBuffer(image.Point{e.WidthPx, e.HeightPx})

				if err != nil {
					log.Fatalf("couldn't create new buffer at size.Event - %v", err)
				}
				// we need to get a new pixel buffer here so we get a buffer of
				// the right size
				pixBuffer = screenBuffer.RGBA()

			case key.Event:
				if e.Code == key.CodeEscape {
					return // quit app
				}
				handleKeyEvents(e)

			case mouse.Event:
				handleMouseEvents(e)

			case paint.Event:
				// fill the background, comment one or the other out below to see
				// the difference when we don't use our updated window sizes.
				// w.Fill(image.Rect(0, 0, 800, 650), color.Black, screen.Src)
				w.Fill(sizeEvent.Bounds(), color.Black, screen.Src)

				// ---- START DRAWING CODE -----
				// 256 pixel square at 100 down and 100 across from window edge
				// using the 'Set' method, limited colors only.
				x_start, x_finish, y_start, y_finish := 100, 256, 100, 256
				for x := x_start; x < x_finish; x++ {
					for y := y_start; y < y_finish; y++ {
						pixBuffer.Set(x, y, color.White)
					}
				}

				// 256 pixel square at 100 down and 300 across from window edge
				// using the 'SetRGBA' method, colors calculated in loop.
				// NOTE: see https://lodev.org/cgtutor/quickcg.html
				x_start, x_finish, y_start, y_finish = 300, 556, 100, 356
				var ciX, ciY uint8 // color index counter
				for x := x_start; x < x_finish; x++ {
					for y := y_start; y < y_finish; y++ {
						pixBuffer.SetRGBA(x, y, color.RGBA{ciX, ciY, uint8(128), 0xff})

						ciY++
					}
					ciX++
				}

				// now let's draw a red line using SetRGBA hex values.
				for x := 400; x < 550; x++ {
					pixBuffer.SetRGBA(x, x, color.RGBA{0xff, 0x00, 0x00, 0xff})
				}
				// ---- END DRAWING CODE ----------

				// upload the pixBuffer (or SwapBuffers if you like)
				w.Upload(image.Point{0, 0}, screenBuffer, sizeEvent.Bounds())
				// finfished drawing etc, swap back buffer to front:
				w.Publish()

			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					// Do any final cleanup or saving here:
					return // quit the application.
				}

			}
		}
	})
}

func handleKeyEvents(e key.Event) {
	switch e.Code {

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
	}
}

func handleMouseEvents(e mouse.Event) {
	switch e.Button {
	case mouse.ButtonLeft:
		if e.Direction == mouse.DirPress {
			fmt.Println("left mouse button down")
		}

		if e.Direction == mouse.DirNone {
			fmt.Println("left mouse button being held")
		}
		if e.Direction == mouse.DirRelease {
			fmt.Println("left mouse button released")
		}
	}
}
