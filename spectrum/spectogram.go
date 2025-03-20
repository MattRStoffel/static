package spectrum

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"os/exec"
)

func fft() {

}

func GenerateSpectrogram(path string, start int, duration int) ([][]uint8, error) {
	formatSeconds := func(seconds int) string {
		return fmt.Sprintf("%02d:%02d:%02d",
			(seconds / 3600),
			(seconds/60)%60,
			(seconds % 60),
		)
	}
	cmd := exec.Command(
		"ffmpeg",
		"-ss", formatSeconds(start),
		"-t", formatSeconds(duration),
		"-i", path,
		"-filter_complex",
		"[0]showspectrumpic=s=600x200:mode=combined:legend=0:start=200:stop=9000",
		"-c:v", "png",
		"-f", "image2pipe",
		"pipe:1",
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	imgBytes, err := io.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	img, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	matrix := make([][]uint8, height)
	for y := 0; y < height; y++ {
		matrix[y] = make([]uint8, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((r + g + b) / 3 >> 8)
			matrix[y][x] = gray
		}
	}

	return matrix, nil
}
