package psdui

import (
	"fmt"

	"github.com/oov/psd"
)

// ImageParser 图片解析类
type ImageParser struct {
	WidgetParser
}

// Parse 解析图层
func (p *ImageParser) Parse(node *UINode, layer *psd.Layer) *UIAst {
	fmt.Println("ImageParser:", layer.Name)
	return nil
}

func init() {
	RegisterParser("Image", new(ImageParser))
}
