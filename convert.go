package imgfilter

import (
	"fmt"
	"math"
	"os/exec"
)

func ExecConvert(args []string) ([]byte, error) {
	cmd := exec.Command("convert", args...)
	return cmd.Output()
}

func ColorTone(inputFile, outputFile, color string, level, ntype int) error {
	colorArg := fmt.Sprintf("( -clone 0 -fill '%s' -colorize 100% )", color)
	negate := ""
	if ntype == 0 {
		negate = "-negate"
	}
	colorSpaceArg := fmt.Sprintf("( -clone 0 -colorspace gray %s )", negate)
	composeArgs := fmt.Sprintf("compose:args=%d,%d", level, 100-level)
	args := []string{inputFile, colorArg, colorSpaceArg, "-compose", "blend", "-define", composeArgs, "-composite", outputFile}
	_, err := ExecConvert(args)
	return err
}

func Border(inputFile, outputFile, color string, width int) error {
	borderArg := fmt.Sprintf("%dx%d", width, width)
	args := []string{inputFile, "-bordercolor", color, "-border", borderArg, inputFile}
	_, err := ExecConvert(args)
	return err
}

func Frame(inputFile, outputFile, frame string, width, height int) error {
	resizeArg := fmt.Sprintf("%dx%d!", width, height)
	args := []string{inputFile, "(", fmt.Sprintf("'%s'", frame), "-resize", resizeArg, "-unsharp", "1.5x1.0+1.5+0.02}", ")", "-flatten", outputFile}
	_, err2 := ExecConvert(args)
	return err2
}

func Vignette(inputFile, outputFile, color1, color2 string, width, height int, cropFactor float32) error {
	cropX := math.Floor(float64(cropFactor * float32(width)))
	cropY := math.Floor(float64(cropFactor * float32(height)))
	sizeArg := fmt.Sprintf("%fx%f", cropX, cropY)
	radialGradientArg := fmt.Sprintf("radial-gradient:%s-%s", color1, color2)
	cropArg := fmt.Sprintf("%dx%d+0+0")
	args := []string{"(", inputFile, ")", "(", "-size", sizeArg, radialGradientArg, "-gravity", "center", "-crop", cropArg, "+repage", ")", "-compose", "multiply", "-flatten", outputFile}
	_, err := ExecConvert(args)
	return err
}
