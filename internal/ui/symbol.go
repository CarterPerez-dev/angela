// ©AngelaMos | 2026
// symbol.go

package ui

import "strings"

// Unicode symbols for CLI output formatting.
const (
	Arrow       = "→"
	ArrowRight  = "▸"
	ArrowUp     = "↑"
	Diamond     = "◆"
	Gem         = "◈"
	Star        = "✦"
	TriangleUp  = "▲"
	Check       = "✓"
	Cross       = "✗"
	Timer       = "⏱"
	DividerChar = "━"
)

// HRule returns a horizontal rule of the given width.
func HRule(width int) string {
	return strings.Repeat(DividerChar, width)
}
