package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
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

	delta float64 = 0.0 // time delta
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

		// Create the starfield object
		stars := Stars3d{
			numStars: 4096,
			spread:   64.0,
			speed:    20.0,
		}
		// we have to init to set the array sizes etc.
		stars.InitStars()

		//var frames = 0
		var previousTime = time.Now()
		for {
			// render the star field
			stars.UpdateAndRender(w, pixBuffer, delta)

			w.Upload(image.Point{0, 0}, screenBuffer, screenBuffer.Bounds())
			w.Publish()
			time.Sleep(time.Millisecond * 6) // slow it down a bit

			// print out the ms/frame value
			currTime := time.Now()
			delta = float64(currTime.Sub(previousTime).Nanoseconds()) / 1000000000.0
			fmt.Printf("\rRendering at %.5f \tms/frame\t", delta*1000)
			previousTime = currTime

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

// clearBuffer - clears the screen to a standard color, can be changed
// to take an RGBA and a flag for gradient perhaps.
func clearBuffer(c color.RGBA) {
	for x := 0; x < pixBuffer.Bounds().Dx(); x++ {
		for y := 0; y < pixBuffer.Bounds().Dy(); y++ {
			pixBuffer.Set(x, y, c)
		}
	}
}

type Stars3d struct {
	numStars int
	spread   float64
	speed    float64
	starX    []float64
	starY    []float64
	starZ    []float64
}

// UpdateAndRender - draws a field of stars to the pixel buffer.
func (s *Stars3d) UpdateAndRender(w screen.Window, img *image.RGBA, delta float64) {
	clearBuffer(color.RGBA{0x00, 0x00, 0x00, 0xff})

	// draw the stars
	halfWidth := float64(pixBuffer.Bounds().Dx()) / 2.0
	halfHeight := float64(pixBuffer.Bounds().Dy()) / 2.0
	for i := 0; i < s.numStars; i++ {
		s.starZ[i] -= delta * s.speed
		if s.starZ[i] <= 0 {
			s.initStar(i)
		}
		x := int((s.starX[i]/s.starZ[i])*halfWidth + halfWidth)
		y := int((s.starY[i]/s.starZ[i])*halfHeight + halfHeight)

		if (x < 0 || x >= pixBuffer.Bounds().Dx()) ||
			(y < 0 || y >= pixBuffer.Bounds().Dy()) {
			s.initStar(i)
		} else {
			pixBuffer.Set(x, y, color.White)
		}
	}
	// force the window to paint itself.
	w.Send(paint.Event{External: true})
}

// InitStars - set up the starfield
func (s *Stars3d) InitStars() {
	s.starX = make([]float64, s.numStars)
	s.starY = make([]float64, s.numStars)
	s.starZ = make([]float64, s.numStars)

	for i := 0; i < s.numStars; i++ {
		s.initStar(i)
	}
}

func (s *Stars3d) initStar(i int) {
	s.starX[i] = (2*rand.Float64() - 0.5) * s.spread
	s.starY[i] = (2*rand.Float64() - 0.5) * s.spread
	s.starZ[i] = (rand.Float64() + 0.00001) * s.spread
}
