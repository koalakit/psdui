package psdui

import (
	"fmt"
	"runtime"
)

// VersionBinary 版本号
const VersionBinary = "0.1.0-alpha"

// VersionString 返回版本号
func VersionString(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, VersionBinary, runtime.Version())
}
