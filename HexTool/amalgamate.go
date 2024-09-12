package HexTool

import "strings"

// HexProcessForMitmWeb
// 针对Mitmweb粘贴的处理，前后无用内容，去除行，保留字节间空格。
// 例如：00000000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  |................|
// 处理后：00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
func HexProcessForMitmWeb(raw string) string {
	s := ""
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		s = s + " " + line[11:58]
	}
	return strings.TrimLeft(s, " ")
}

// CleanSpace 清理掉所有字符串间空格
func CleanSpace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
