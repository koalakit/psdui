package psdui

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ExportPreview 输出预览图
func ExportPreview(sourceFile string, outputDir string) error {
	file, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	filename := filepath.Base(sourceFile)
	filename = strings.Replace(filename, ".psd", ".png", -1)

	if len(outputDir) <= 0 {
		outputDir = filepath.Dir(sourceFile)
	}

	_, err = os.Stat(outputDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			return err
		}
	}

	out, err := os.Create(filepath.Join(outputDir, filename))
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}
