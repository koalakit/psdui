package egret

import (
	"encoding/json"
	"fmt"

	"github.com/koalakit/psdui"
)

// UIDIR UI目录
const UIDIR = "assets/ui"

// EXMLDIR exml文件目录
const EXMLDIR = "assets/skins"

// EgretExpoter Egret 导出工具
type EgretExpoter struct {
}

// Export 导出资源
func (exporter *EgretExpoter) Export(node *psdui.UINode, outputDir string) error {
	nodeBytes, _ := json.Marshal(node)
	fmt.Printf("%s\n", string(nodeBytes))

	return nil
}

func init() {
	psdui.AddExporter("egret", new(EgretExpoter))
}
