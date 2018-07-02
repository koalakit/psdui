package psdui

import (
	"fmt"

	"github.com/oov/psd"
)

type ComponentParser struct {
}

// Parse 解析图层
func (p *ComponentParser) Parse(node *UINode, layer *psd.Layer) *UIAst {
	fmt.Println("ComponentParser:", layer.Name)

	// var ui UIAst
	// prop.Attributes["位置"]
	return nil
}

func init() {
	RegisterParser("UI", new(ComponentParser))
}
