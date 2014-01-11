package imgfilter

import (
	"fmt"
	"path"
)

func outputFile(suffix, inputFile, outputDir string) string {
	inputFileBase := path.Base(inputFile)
	ext := path.Ext(inputFileBase)
	outputFileBase := fmt.Sprintf("%s_%s%s", inputFileBase[0:len(inputFileBase)-len(ext)], suffix, ext)
	return path.Join(outputDir, outputFileBase)
}

func FilterGothem(inputFile, outputDir string) error {
	convert := []string{inputFile, "-modulate", "120,10,100", "-fill", "#222b6d", "-colorize", "20", "-gamma", "0.5", "-contrast", "-contrast", outputFile("gothem", inputFile, outputDir)}
	_, err := ExecConvert(convert)
	return err
}
