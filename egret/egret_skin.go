package egret

import (
	"fmt"
	"image"
	"os"

	"github.com/beevik/etree"
	"github.com/koalakit/psdui"
)

var (
	// ErrRootNodeNameInvalid 根节点类型无效
	ErrRootNodeNameInvalid = fmt.Errorf(`根节点类型无效，必须为"UI"`)
)

// EgretSkin 皮肤节点
type EgretSkin struct {
	Class   string `xml:"class,attr"`
	X       int
	Y       int
	Width   int    `xml:"width,attr"`
	Height  int    `xml:"height,attr"`
	NS      string `xml:""`
	Rect    image.Rectangle
	XMLDoc  *etree.Document
	XMLNode *etree.Element
}

// Parse 解析节点
func (skin *EgretSkin) Parse(node *psdui.UINode) error {
	if node.Type != "UI" {
		return ErrRootNodeNameInvalid
	}

	skin.XMLDoc = etree.NewDocument()
	skin.XMLDoc.CreateProcInst("xml", `version="1.0" encoding="utf-8"`)

	skinNode := skin.XMLDoc.CreateElement("e:Skin")
	skinNode.CreateAttr("xmlns:e", "http://ns.egret.com/eui")
	skin.XMLNode = skinNode

	// 解析坐标属性
	var positionAttr PositionAttribute
	if node.Attr(&positionAttr) {
		skin.X = positionAttr.X
		skin.Y = positionAttr.Y

		positionAttr.Write(skinNode)
	}

	// 解析尺寸属性
	var sizeAttr SizeAttribute
	if node.Attr(&sizeAttr) {
		skin.Width = sizeAttr.Width
		skin.Height = sizeAttr.Height

		sizeAttr.Write(skinNode)
	}

	skin.XMLDoc.Indent(4)
	skin.XMLDoc.WriteTo(os.Stdout)

	return nil
}

var skinRoot = new(EgretSkin)

////////////////////////////////////////////////////////////////////////////////
// 最小尺寸属性

// MinSizeAttribute 最小尺寸属性
type MinSizeAttribute struct {
	Width  int `json:"width,omitempy"`
	Height int `json:"height,omitempy"`
}

// Name 返回名称
func (attr *MinSizeAttribute) Name() string {
	return MinSizeName
}

// Parse 解析属性
func (attr *MinSizeAttribute) Parse(v string) error {
	n, err := fmt.Sscanf(v, "%d,%d", &attr.Width, &attr.Height)
	if err != nil {
		return err
	}

	if n != 2 {
		return fmt.Errorf("最小尺寸属性格式错误:%s 正确格式: (width,height)", v)
	}

	return nil
}

// Write 输出到xml
func (attr *MinSizeAttribute) Write(node *etree.Element) {
	node.CreateAttr("minWidth", fmt.Sprint(attr.Width))
	node.CreateAttr("minHeight", fmt.Sprint(attr.Height))
}

func init() {
	psdui.RegisterAttribute(new(MinSizeAttribute))
}
