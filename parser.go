package psdui

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/oov/psd"
)

// PSDParser PSD 文件解析
type PSDParser struct {
	SourceFile string
	ImagePSD   *psd.PSD
	Width      int
	Height     int
	Nodes      []*UINode
}

// Load 载入PSD文件
func (parser *PSDParser) Load(sourceFile string) (err error) {
	parser.SourceFile = sourceFile

	var file *os.File
	file, err = os.Open(sourceFile)
	if err != nil {
		return
	}
	defer file.Close()

	// var options psd.DecodeOptions
	parser.ImagePSD, _, err = psd.Decode(file, nil)
	if err != nil {
		return
	}

	config := parser.ImagePSD.Config
	parser.Width = config.Rect.Dx()
	parser.Height = config.Rect.Dy()

	fullpath, _ := filepath.Abs(sourceFile)
	fmt.Printf("%s(%d:%d)\n", fullpath, parser.Width, parser.Height)

	err = parser.parse()
	return
}

// 解析图层
func (parser *PSDParser) parse() (err error) {
	parser.Nodes = make([]*UINode, 0)
	for _, v := range parser.ImagePSD.Layer {
		layer := v
		node := parser.processLayer(nil, &layer)
		if node != nil {
			parser.Nodes = append(parser.Nodes, node)
		}
	}

	return err
}

func (parser *PSDParser) processLayer(parent *UINode, layer *psd.Layer) *UINode {
	name := layer.Name
	node := ParseNode(name)

	var subParent *UINode
	subParent = parent

	if node != nil {
		node.SetLayer(layer)
		node.ImagePSD = parser.ImagePSD
		subParent = node
	}

	for _, v := range layer.Layer {
		subLayer := v
		parser.processLayer(subParent, &subLayer)
	}

	if node != nil {
		if parent != nil {
			parent.AddChild(*node)
		}
	}

	return node
}

// NewPSDParser 分析PSD解析器
func NewPSDParser() *PSDParser {
	parser := new(PSDParser)
	return parser
}

// // IParser 解析器接口类
// type IParser interface {
// 	Parse(node *UINode, layer *psd.Layer) *UIAst
// }

// // WidgetParser 控件解析基类
// type WidgetParser struct {
// 	X      int
// 	Y      int
// 	Width  int
// 	Height int
// }

// // Parse 解析图层
// func (w *WidgetParser) Parse(node *UINode, layer *psd.Layer) error {
// 	return nil
// }

// var (
// 	_parserFactories = make(map[string]IParser)
// )

// // RegisterParser 注册解析类
// func RegisterParser(name string, parser IParser) {
// 	_parserFactories[name] = parser
// }

// // NewWidgetParser 分配新控件解析器
// func NewWidgetParser(typ string) IParser {
// 	parser, ok := _parserFactories[typ]
// 	if !ok {
// 		return nil
// 	}

// 	return parser
// }
