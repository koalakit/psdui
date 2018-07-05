package egret

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/koalakit/psdui"
)

////////////////////////////////////////////////////////////////////////////////
// 属性名称
const (
	// PositionName 位置属性名称
	PositionName = "位置"
	// SizeName 尺寸属性名称
	SizeName = "尺寸"
	// TextName 文本属性名称
	TextName = "文本"
	// ExportName 导出属性名称
	ExportName = "导出"
	// Scale9GridName 九宫格属性名称
	Scale9GridName = "九宫格"

	// MinSizeName 最小尺寸属性
	MinSizeName = "最小尺寸"
)

// IAttributeXML 属性输出到xml
type IAttributeXML interface {
	Write(node *etree.Element)
}

////////////////////////////////////////////////////////////////////////////////
// 位置属性

// PositionAttribute 位置属性
type PositionAttribute struct {
	X int `json:"x,omitempy"`
	Y int `json:"y,omitempy"`
}

// Name 返回名称
func (attr *PositionAttribute) Name() string {
	return PositionName
}

// Parse 解析属性
func (attr *PositionAttribute) Parse(v string) error {
	n, err := fmt.Sscanf(v, "%d,%d", &attr.X, &attr.Y)
	if err != nil {
		return err
	}

	if n != 2 {
		return fmt.Errorf("位置属性格式错误:%s 正确格式: (x,y)", v)
	}

	return nil
}

// Write 输出到xml
func (attr *PositionAttribute) Write(node *etree.Element) {
	node.CreateAttr("x", fmt.Sprint(attr.X))
	node.CreateAttr("y", fmt.Sprint(attr.Y))
}

////////////////////////////////////////////////////////////////////////////////
// 尺寸属性

// SizeAttribute 尺寸属性
type SizeAttribute struct {
	Width  int `json:"width,omitempy"`
	Height int `json:"height,omitempy"`
}

// Name 返回名称
func (attr *SizeAttribute) Name() string {
	return SizeName
}

// Parse 解析属性
func (attr *SizeAttribute) Parse(v string) error {
	n, err := fmt.Sscanf(v, "%d,%d", &attr.Width, &attr.Height)
	if err != nil {
		return err
	}

	if n != 2 {
		return fmt.Errorf("尺寸属性格式错误:%s 正确格式: (width,height)", v)
	}

	return nil
}

// Write 输出到xml
func (attr *SizeAttribute) Write(node *etree.Element) {
	node.CreateAttr("width", fmt.Sprint(attr.Width))
	node.CreateAttr("height", fmt.Sprint(attr.Height))
}

////////////////////////////////////////////////////////////////////////////////
// 文本属性

// TextAttribute 文本属性
type TextAttribute struct {
	Text string `json:"text,omitempy"`
}

// Name 返回名称
func (attr *TextAttribute) Name() string {
	return TextName
}

// Parse 解析属性
func (attr *TextAttribute) Parse(v string) error {
	attr.Text = v

	return nil
}

// Write 输出到xml
func (attr *TextAttribute) Write(node *etree.Element) {
	node.CreateAttr("text", attr.Text)
}

////////////////////////////////////////////////////////////////////////////////
// 导出属性

// ExportAttribute 导出属性
type ExportAttribute struct {
	ExportName string `json:"export,omitempy"`
}

// Name 返回名称
func (attr *ExportAttribute) Name() string {
	return ExportName
}

// Parse 解析属性
func (attr *ExportAttribute) Parse(v string) error {
	attr.ExportName = v

	return nil
}

// Write 输出到xml
func (attr *ExportAttribute) Write(node *etree.Element) {
	// node.CreateAttr("export", attr.ExportName)
}

////////////////////////////////////////////////////////////////////////////////
// 九宫格属性

// Scale9GridAttribute 九宫格属性
type Scale9GridAttribute struct {
	X      int `json:"x,omitempy"`
	Y      int `json:"y,omitempy"`
	Width  int `json:"width,omitempy"`
	Height int `json:"height,omitempy"`
}

// Name 返回名称
func (attr *Scale9GridAttribute) Name() string {
	return Scale9GridName
}

// Parse 解析属性
func (attr *Scale9GridAttribute) Parse(v string) error {
	n, err := fmt.Sscanf(v, "%d,%d,%d,%d", &attr.X, &attr.Y, &attr.Width, &attr.Height)
	if err != nil {
		return err
	}

	if n != 4 {
		return fmt.Errorf("九宫格属性格式错误:%s 正确格式: (x,y,width,height)", v)
	}

	return nil
}

// Write 输出到xml
func (attr *Scale9GridAttribute) Write(node *etree.Element) {
	node.CreateAttr("scale9Grid", fmt.Sprintf("%d,%d,%d,%d", attr.X, attr.Y, attr.Width, attr.Height))
}

// 注册属性
func init() {
	psdui.RegisterAttribute(new(PositionAttribute))
	psdui.RegisterAttribute(new(SizeAttribute))
	psdui.RegisterAttribute(new(TextAttribute))
	psdui.RegisterAttribute(new(ExportAttribute))
	psdui.RegisterAttribute(new(Scale9GridAttribute))
}
