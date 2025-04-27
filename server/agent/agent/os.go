package agent

import (
	"runtime"
)

// GetOSType 返回当前操作系统类型，windows/linux/macos
func GetOSType() string {
	switch runtime.GOOS {
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	case "darwin": // macOS 在 Go 中是 "darwin"
		return "macos"
	default:
		return "unknown" // 如果不是 windows、linux 或 macos，返回 unknown
	}
}
