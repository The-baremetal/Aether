package utils

import (
	"path/filepath"
)

// ModulePathToFilePath converts a module path to a file path.
// Example: "module/package" -> "module/package/module.ae"
func ModulePathToFilePath(modulePath string) string {
	return filepath.Join(modulePath, "module.ae")
}
