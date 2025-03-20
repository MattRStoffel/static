package main

import (
	"fmt"
	"static/fingerprint"
	"static/spectrum"
)

func main() {
	sepctrum, err := spectrum.GenerateSpectrogram("test_data/input2.mp3", 3, 3)
	if err != nil {
		fmt.Println("ERRORR: ", err)
	}
	brightPoints := fingerprint.Fingerprint(sepctrum, 300)
	for _, point := range brightPoints {
		fmt.Println("X:", point.X, "Y:", point.Y, "Intensity:", point.Intensity)
	}
}
