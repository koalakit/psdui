package psdui

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/oov/psd"
)

// ExportPreview 输出预览图
func ExportPreview(sourceFile string, outputDir string) error {
	file, err := os.Open(sourceFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	filename := filepath.Base(sourceFile)
	filename = strings.Replace(filename, ".psd", ".png", -1)
	_, err = os.Stat(outputDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	out, err := os.Create(filepath.Join(outputDir, filename))
	if err != nil {
		panic(err)
	}

	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}
	return nil
}
