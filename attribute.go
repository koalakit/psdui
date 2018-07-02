package psdui

import (
	"fmt"
	"reflect"
)

// IAttribute 属性接口
type IAttribute interface {
	Parse(v string) error
}

// PositionAttribute 位置属性
type PositionAttribute struct {
	X int `json:"x,omitempy"`
	Y int `json:"y,omitempy"`
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

// SizeAttribute 位置属性
type SizeAttribute struct {
	Width  int `json:"width,omitempy"`
	Height int `json:"height,omitempy"`
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

// TextAttribute 文本属性
type TextAttribute struct {
	Text string `json:"text,omitempy"`
}

// Parse 解析属性
func (attr *TextAttribute) Parse(v string) error {
	attr.Text = v

	return nil
}

// Scale9GridAttribute 位置属性
type Scale9GridAttribute struct {
	X      int `json:"x,omitempy"`
	Y      int `json:"y,omitempy"`
	Width  int `json:"width,omitempy"`
	Height int `json:"height,omitempy"`
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

var (
	_attributeFactories = make(map[string]reflect.Type)
)

// RegisterAttribute 注册属性
func RegisterAttribute(name string, attr IAttribute) {
	_attributeFactories[name] = reflect.ValueOf(attr).Elem().Type()
}

// NewAttribute 通过名称分配新属性
func NewAttribute(name string) IAttribute {
	attrType, ok := _attributeFactories[name]
	if !ok {
		return nil
	}

	attrValue := reflect.New(attrType)
	attr, ok := attrValue.Interface().(IAttribute)
	if !ok {
		return nil
	}

	return attr
}

func init() {
	RegisterAttribute("位置", new(PositionAttribute))
	RegisterAttribute("尺寸", new(SizeAttribute))
	RegisterAttribute("文本", new(TextAttribute))
	RegisterAttribute("九宫格", new(Scale9GridAttribute))
}
