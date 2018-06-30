package psdui

import (
	"fmt"
	"os"

	"github.com/oov/psd"
)

// PsdFile PSD文件操作
type PsdFile struct {
	psdImage *psd.PSD
	Width    int
	Height   int
}

func (f *PsdFile) load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	f.psdImage, _, err = psd.Decode(file, &psd.DecodeOptions{
		SkipMergedImage: true,
	})

	if err != nil {
		return err
	}

	bounds := f.psdImage.Picker.Bounds()
	f.Width = bounds.Max.X - bounds.Min.X
	f.Height = bounds.Max.Y - bounds.Min.Y

	return nil
}

func (f *PsdFile) String() string {
	return fmt.Sprintf("PSD %d:%d", f.Width, f.Height)
}
