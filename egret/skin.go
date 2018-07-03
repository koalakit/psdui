package egret

import (
	"fmt"
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
	xmlDoc  *etree.Document
	xmlNode *etree.Element
}

// Parse 解析节点
func (skin *EgretSkin) Parse(node *psdui.UINode) error {
	if node.Type != "UI" {
		return ErrRootNodeNameInvalid
	}

	skin.xmlDoc = etree.NewDocument()
	skin.xmlDoc.CreateProcInst("xml", `version="1.0" encoding="utf-8"`)

	skinNode := skin.xmlDoc.CreateElement("e:Skin")
	skinNode.CreateAttr("xmlns:e", "http://ns.egret.com/eui")
	skin.xmlNode = skinNode

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

	skin.xmlDoc.Indent(4)
	skin.xmlDoc.WriteTo(os.Stdout)

	return nil
}

var skinRoot = new(EgretSkin)
