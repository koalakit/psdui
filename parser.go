package psdui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/oov/psd"
)

// PSDParser PSD 文件解析
type PSDParser struct {
	SourceFile string
	ImagePSD   *psd.PSD
	Width      int
	Height     int
}

// Load 载入PSD文件
func (parser *PSDParser) Load(sourceFile string) error {
	parser.SourceFile = sourceFile

	var err error
	var file *os.File
	file, err = os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer file.Close()

	parser.ImagePSD, _, err = psd.Decode(file, &psd.DecodeOptions{
		SkipMergedImage: true,
	})
	if err != nil {
		return err
	}

	config := parser.ImagePSD.Config
	parser.Width = config.Rect.Dx()
	parser.Height = config.Rect.Dy()

	fullpath, _ := filepath.Abs(sourceFile)
	fmt.Printf("%s(%d:%d)\n", fullpath, parser.Width, parser.Height)

	return parser.parse()
}

func (parser *PSDParser) parse() error {
	var err error

	for _, layer := range parser.ImagePSD.Layer {
		if err = parser.processLayer(&layer); err != nil {
			panic(err)
		}
	}

	return err
}

func (parser *PSDParser) processLayer(layer *psd.Layer) error {
	name := layer.Name
	parser.parseName(name)

	var err error

	if len(layer.Layer) > 0 {
		for _, subLayer := range layer.Layer {
			if err = parser.processLayer(&subLayer); err != nil {
				return err
			}
		}
	}

	return nil
}

func (parser *PSDParser) parseName(name string) bool {
	if len(name) <= 0 {
		return false
	}

	typeIdx := strings.Index(name, "@")
	if typeIdx < 0 {
		return false
	}

	if typeIdx+1 >= len(name) {
		fmt.Printf("[ERROR] 图层[%v] 字符@后必须定义类型\n", name)
		return false
	}

	elementName := name[:typeIdx]
	typeIdx++
	elementType := name[typeIdx:]
	attrIdx := strings.Index(elementType, ":")
	if attrIdx == 0 {
		fmt.Printf("[ERROR] 图层[%v] 字符@后必须定义类型\n", name)
		return false
	} else if attrIdx < 0 {
		elementType = name[typeIdx:]
	} else {
		elementType = name[typeIdx : typeIdx+attrIdx]
	}

	fmt.Println("name:", elementName, "type:", elementType)

	return true
}

// NewPSDParser 分析PSD解析器
func NewPSDParser() *PSDParser {
	parser := new(PSDParser)
	return parser
}
