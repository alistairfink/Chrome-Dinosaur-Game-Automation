package main

import (
	"github.com/go-vgo/robotgo"
	"time"
	// "github.com/vova616/screenshot"
)

func main() {
	println("Before starting click on the dinosaur's eye...")
	mleft := robotgo.AddEvent("mleft")
	var mouseX, mouseY int
	if mleft {
		mouseX, mouseY = robotgo.GetMousePos()
	}

	pixelWallX, pixelWallYOffset := 80, 100-10
	noObstacleScreenshot := robotgo.CaptureScreen(mouseX, mouseY-100, 100, 150)
	noObstacleColour := robotgo.GetColor(noObstacleScreenshot, pixelWallX, pixelWallYOffset)
	robotgo.FreeBitmap(noObstacleScreenshot)

	println("Starting...")
	robotgo.KeyTap("space")

	i := 1
	for {
		screenShot := robotgo.CaptureScreen(mouseX, mouseY-100, 100, 150)
		y := 36
		verticalThresh := 8
		for ; y >= verticalThresh; y -= 2 {
			curr := robotgo.GetColor(screenShot, pixelWallX, pixelWallYOffset+y)
			if curr != noObstacleColour {
				jump(&i)
				break
			}
		}

		if y <= verticalThresh {
			for ; y >= 0; y -= 2 {
				curr := robotgo.GetColor(screenShot, pixelWallX+20, pixelWallYOffset+y)
				if curr != noObstacleColour {
					jump(&i)
					break
				}
			}
		}

		robotgo.FreeBitmap(screenShot)
	}
}

func jump(i *int) {
	robotgo.KeyTap("space")
	println("jump", *i)
	*i++
	time.Sleep(400 * time.Millisecond)
}
