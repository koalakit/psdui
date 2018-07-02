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

// NameProperty 名称属性
type NameProperty struct {
	Name       string
	Type       string
	Attributes map[string]string
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

// 解析图层
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
	_, ok := parser.parseName(name)
	if ok {
		// fmt.Printf("LAYER: %s %v:%v\n", name, layer.Rect.Dx(), layer.Rect.Dy())
	}

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

// 如果不包含@字符则返回false，直接跳过
func (parser *PSDParser) parseName(name string) (property NameProperty, ok bool) {
	ok = false

	if len(name) <= 0 {
		return
	}

	// 解析名称类型
	typeIdx := strings.Index(name, "@")
	if typeIdx < 0 {
		return
	}

	if typeIdx+1 >= len(name) {
		fmt.Printf("[ERROR] 图层[%v] 字符@后必须定义类型\n", name)
		return
	}

	elementName := name[:typeIdx]
	typeIdx++
	elementType := name[typeIdx:]
	var elementAttrs string

	attrIdx := strings.Index(elementType, ":")
	if attrIdx == 0 {
		fmt.Printf("[ERROR] 图层[%v] 字符@后必须定义类型\n", name)
		return
	} else if attrIdx < 0 {
		elementType = name[typeIdx:]
	} else {
		elementType = name[typeIdx : typeIdx+attrIdx]
		elementAttrs = name[typeIdx+attrIdx:]
	}

	property.Name = elementName
	property.Type = elementType

	// 属性定义是否合法
	if len(elementAttrs) == 1 {
		fmt.Printf("[ERROR] 图层[%v] 类型后需要添加属性，如没有自定义属性请去掉类型后的“:”符号\n", name)
		return
	}

	// 解析属性
	if len(elementAttrs) > 0 {
		elementAttrs = elementAttrs[1:]
	}

	ok = true

	// 解析属性
	property.Attributes = parser.parseAttributes(name, elementAttrs)

	fmt.Printf("%+v\n", property)
	return
}

func (parser *PSDParser) parseAttributes(name string, str string) map[string]string {
	if len(str) <= 0 {
		return nil
	}

	attributes := make(map[string]string)

	str1 := strings.Split(str, ";")
	leftBracket := 0
	rightBracket := 0
	for _, v := range str1 {
		leftBracket = strings.Index(v, "(")
		if leftBracket < 0 {
			// 有默认值
			attributes[v] = ""
			continue
		}

		rightBracket = strings.Index(v, ")")
		if rightBracket < 0 {
			fmt.Printf("[ERROR] 图层[%v] 属性值的括号必须匹配\n", name)
			continue
		}

		attributes[v[:leftBracket]] = v[leftBracket+1 : rightBracket]
	}

	return attributes
}

// NewPSDParser 分析PSD解析器
func NewPSDParser() *PSDParser {
	parser := new(PSDParser)
	return parser
}
