package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"golang.org/x/mobile/event/lifecycle"

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
			drawScene(w)
			w.Upload(image.Point{0, 0}, screenBuffer, screenBuffer.Bounds())
			w.Publish()
			time.Sleep(time.Millisecond * 16)

			// print out the ms/frame value
			frames++
			currTime = time.Now()
			if currTime.Sub(startTime).Seconds() >= 1.0 {
				fmt.Printf("\rRendering at %.5f \tms/frame", 1000.0/float64(frames))
				frames = 0
				startTime = currTime
			}

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

func drawScene(w screen.Window) {
	// load the DOOM image:
	// Read image from file that already exists
	existingImageFile, err := os.Open("doom.png")
	if err != nil {
		log.Fatalf("%v - failed to open image", err)
	}
	defer existingImageFile.Close()
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		log.Fatalf("%v - failed to decode image", err)
	}

	// calculate the values to set the image to the center of the screen buffer.
	// NOTE: this doesn't seem right and needs more investigation!
	// I think it has a bit to do with the actual image size compared to the
	// actual image displayed but still not sure...
	imgRect := loadedImage.Bounds()
	imgX := ((pixBuffer.Rect.Dx() - imgRect.Dx()) / 2) - imgRect.Dx()
	imgY := ((pixBuffer.Rect.Dy() - imgRect.Dy()) / 2) - imgRect.Dy()

	// to draw the image we need to 'Draw' it to the pixel buffer.
	draw.Draw(pixBuffer, pixBuffer.Bounds(), loadedImage, image.Point{imgX, imgY}, draw.Src)

	// now let's draw a red line using SetRGBA hex values.
	// this line should draw over the pixels in the previous loaded image.
	for x := 50; x < 550; x++ {
		pixBuffer.SetRGBA(x, x, color.RGBA{0xff, 0x00, 0x00, 0xff})
	}
}
