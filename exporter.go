package psdui

import (
	"fmt"
)

type Exporter interface {
	Export(node *UINode, outputDir string) error
}

var (
	_exporters = make(map[string]Exporter)
)

func AddExporter(name string, exporter Exporter) {
	_exporters[name] = exporter
}

func Export(node *UINode, name string, outputDir string) error {
	exporter, ok := _exporters[name]
	if !ok {
		return fmt.Errorf("导出工具%s未找到", name)
	}

	return exporter.Export(node, outputDir)
}
