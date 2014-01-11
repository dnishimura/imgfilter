package imgfilter

import (
	"image"
	"os"
)

type ImgFilter struct {
	inputFile string
	outputDir string
}

type FilterFunc func(inputFile, outputFile string, width, height int) error

func (imgf *ImgFilter) Execute(filter FilterFunc) error {
	width, height, err := imageDimensions(imgf.inputFile)
	if err != nil {
		return err
	}
	return filter(imgf.inputFile, imgf.outputDir, width, height)
}

func imageDimensions(imgFile string) (int, int, error) {
	f, err := os.Open(imgFile)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return 0, 0, err
	}
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	return width, height, nil
}
