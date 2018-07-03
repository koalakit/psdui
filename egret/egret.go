package egret

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/koalakit/psdui"
)

// UIDIR UI目录
const UIDIR = "assets/ui"

// EXMLDIR exml文件目录
const EXMLDIR = "assets/skins"

// EXMLEXT exml文件后缀
const EXMLEXT = ".exml"

// EgretExporter Egret 导出工具
type EgretExporter struct {
}

// Export 导出资源
func (exporter *EgretExporter) Export(node *psdui.UINode, outputDir string) error {
	// nodeBytes, _ := json.Marshal(node)
	// fmt.Printf("%s\n", string(nodeBytes))

	var err error
	if err = skinRoot.Parse(node); err != nil {
		fmt.Println(node.SourceName, err)
	}

	exporter.exportImage(node, outputDir)

	// exmlTest := make(map[string]interface{})
	// data, err := ioutil.ReadFile(filepath.Join(outputDir, "skins/TestUI.exml"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(data))
	// err = json.Unmarshal(data, exmlTest)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v", exmlTest)
	return nil
}

func (exporter *EgretExporter) exportImage(node *psdui.UINode, outputDir string) error {
	if node == nil {
		return nil
	}

	layer := node.Layer

	if layer == nil {
		log.Println("layer is empty")
		return nil
	}

	if layer.HasImage() {
		dir := filepath.Join(outputDir, UIDIR)
		os.MkdirAll(dir, 0777)

		imagePath := filepath.Join(dir, node.Name+".png")
		imageFile, err := os.Create(imagePath)

		if err == nil {
			png.Encode(imageFile, layer.Picker)
			imageFile.Close()
		}
	}

	// log.Println("RECT:", node.Name, layer.Rect)

	for _, v := range node.Children {
		exporter.exportImage(&v, outputDir)
	}

	return nil
}

func init() {
	psdui.AddExporter("egret", new(EgretExporter))
}
