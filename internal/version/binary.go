package version

import (
	"fmt"
	"runtime"
)

// Binary 版本号
const Binary = "0.1.0-alpha"

// String 返回版本号
func String(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, Binary, runtime.Version())
}
