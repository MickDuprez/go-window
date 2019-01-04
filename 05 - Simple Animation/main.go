package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"time"

	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"

	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/size"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

var (
	// set up some global helper var's
	winWidth, winHeight = 600, 400

	// We can get info from the event.Size() function along with other
	// helpful functions and data.
	sizeEvent size.Event

	screenSize   = image.Point{winWidth, winHeight}
	screenBuffer screen.Buffer
	pixBuffer    *image.RGBA
	s            screen.Screen
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

		screenBuffer, err = s.NewBuffer(screenSize)
		if err != nil {
			log.Fatalf("%v - failed to create screen buffer", err)
		}
		defer screenBuffer.Release()
		pixBuffer = screenBuffer.RGBA()

		var frames = 0
		var startTime time.Time
		var currTime = time.Now()
		for {
			clearBuffer(color.Black)
			drawToBuffer(w)
			w.Upload(image.Point{0, 0}, screenBuffer, screenBuffer.Bounds())
			w.Publish()
			time.Sleep(time.Millisecond * 5) // slow it down a bit

			// print out the ms/frame value
			frames++
			currTime = time.Now()
			if currTime.Sub(startTime).Seconds() >= 1.0 {
				fmt.Printf("\rRendering at %.3f \tms/frame\t", 1000.0/float64(frames))
				frames = 0
				startTime = currTime
			} //

			// Handle window events:
			switch e := w.NextEvent().(type) {

			case key.Event:
				if e.Code == key.CodeEscape {
					return // quit app
				}

			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					// Do any final cleanup or saving here:
					return // quit the application.
				}
			}
		}
	})
}

var (
	// create the bounds variables
	boxSize = 50
	top     = 0
	btm     = screenSize.Y - boxSize
	left    = 0
	right   = screenSize.X - boxSize

	// we'll be drawing a simple 50x50 pixel square that will go from
	// top/left to approx. btm/right. Once it reaches the bottom it will
	// start at the top at the last X position and continue.
	// We'll try moving at a set distance of pixels per frame and fine tune to suit.
	pX = left // start pixel in X
	pY = top  // start pixel in Y
)

// This is where we draw the pixels of our image, this will be refactored out soon.
func drawToBuffer(w screen.Window) {

	// check starting pixels are within bounds.
	if pX >= right {
		pX = left
	}
	if pY >= btm {
		pY = top
	}

	// draw the box
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			pixBuffer.SetRGBA(pX+x, pY+y, color.RGBA{0xff, 0x00, 0x00, 0xff})
		}
	}
	// move the starting pixel
	pX += 2
	pY += 2

	// force the window to paint itself.
	w.Send(paint.Event{External: true})
}

// clearBuffer - clears the screen to a standard color, can be changed
// to take an RGBA and a flag for gradient perhaps.
func clearBuffer(c color.Color) {
	for x := 0; x < pixBuffer.Bounds().Dx(); x++ {
		for y := 0; y < pixBuffer.Bounds().Dy(); y++ {
			pixBuffer.Set(x, y, c)
		}
	}
}
