package psdui

import (
	"fmt"
)

// Exporter 导出器
type Exporter interface {
	Export(node *UINode, outputDir string) error
}

var (
	_exporters = make(map[string]Exporter)
)

// AddExporter 添加导出器
func AddExporter(name string, exporter Exporter) {
	_exporters[name] = exporter
}

// Export 导出UI数据
func Export(node *UINode, name string, outputDir string) error {
	exporter, ok := _exporters[name]
	if !ok {
		return fmt.Errorf("导出工具%s未找到", name)
	}

	return exporter.Export(node, outputDir)
}
