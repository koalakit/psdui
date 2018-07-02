package psdui

import (
	"fmt"
	"strings"

	"github.com/oov/psd"
)

// UINode UI原始节点
type UINode struct {
	SeqID      int                   `json:"seqid"`
	SourceName string                `json:"-"`
	Name       string                `json:"name"`
	Type       string                `json:"type"`
	Attributes map[string]IAttribute `json:"attributies,omitempy"`
	Parent     *UINode               `json:"-"`
	Children   []UINode              `json:"children,omitempy"`
	Layer      *psd.Layer            `json:"-"`
	ImagePSD   *psd.PSD              `json:"-"`
}

func (node *UINode) parse(name string) (ok bool) {
	ok = false

	if len(name) <= 0 {
		return
	}

	node.SourceName = name

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

	node.Name = elementName
	node.Type = elementType

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
	node.Attributes = node.parseAttributes(name, elementAttrs)

	for n, attr := range node.Attributes {
		fmt.Printf("NAME: %v:%v\n", n, attr)
	}

	return
}

// parseAttributes 解析属性
func (node *UINode) parseAttributes(name string, str string) map[string]IAttribute {
	if len(str) <= 0 {
		return nil
	}

	attributes := make(map[string]IAttribute)

	str1 := strings.Split(str, ";")
	leftBracket := 0
	rightBracket := 0
	for _, v := range str1 {
		leftBracket = strings.Index(v, "(")
		if leftBracket < 0 {
			// 有默认值
			attributes[v] = nil
			continue
		}

		rightBracket = strings.Index(v, ")")
		if rightBracket < 0 {
			fmt.Printf("[ERROR] 图层[%v] 属性值的括号必须匹配\n", name)
			continue
		}

		// attributes[v[:leftBracket]] = v[leftBracket+1 : rightBracket]
		attrName := v[:leftBracket]
		attrValue := v[leftBracket+1 : rightBracket]

		attr := NewAttribute(attrName)
		if attr == nil {
			fmt.Printf("[ERROR] 图层[%v] %s属性未定义\n", name, attrName)
			continue
		}

		if err := attr.Parse(attrValue); err != nil {
			fmt.Printf("[ERROR] 图层[%v] %v\n", name, err)
			continue
		}

		attributes[attrName] = attr
	}

	return attributes
}

func (node *UINode) SetLayer(layer *psd.Layer) {
	node.Layer = layer
	if layer != nil {
		node.SeqID = layer.SeqID
	}
}

func (node *UINode) AddChild(child UINode) {
	node.Children = append(node.Children, child)
	child.Parent = node
}

func (node *UINode) RemoveChild(child UINode) {
	if len(node.Children) <= 0 {
		return
	}

	for i, v := range node.Children {
		if v.SeqID != child.SeqID {
			continue
		}

		node.Children = append(node.Children[:i], node.Children[i+1:]...)
		child.Parent = nil
		break
	}
}

func (node *UINode) RemoveSelf() {
	if node.Parent != nil {
		node.Parent.RemoveChild(*node)
	}
}

// ParseNode 解析名称，如果不包含@字符则返回false，直接跳过
func ParseNode(name string) *UINode {
	var node UINode
	if node.parse(name) {
		return &node
	}

	return nil
}
