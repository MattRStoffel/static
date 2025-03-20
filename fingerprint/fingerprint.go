package fingerprint

import (
	"fmt"
	"sort"
)

type Coords struct {
	X         int
	Y         int
	Intensity int
}

type Song struct {
	Hash string
	Name string
	Time string
}

func Fingerprint(spectrum [][]uint8, numPoints int) []Coords {
	points := []Coords{}

	for y := 0; y < len(spectrum); y++ {
		for x := 0; x < len(spectrum[y]); x++ {
			fmt.Println("HELLO")
			intensity := int(spectrum[y][x])
			points = append(points, Coords{X: x, Y: y, Intensity: int(intensity)})
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].Intensity > points[j].Intensity
	})

	if len(points) > numPoints {
		points = points[:numPoints]
	}

	return points
}
