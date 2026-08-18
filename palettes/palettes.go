// Package palettes provides the official Leraniode colour palettes built on
// wondertone.
//
// Available palettes:
//
//   - Leraniode  — Leraniode full base palette
//   - LeraScale  — Leraniode pink identity scale
//   - NiodeScale — Leraniode purple identity scale
//   - Clarity    — Leraniode cyan-to-blue gradient scale
package palettes

import (
	"github.com/leraniode/wondertone/palette"
)

var registry []*palette.Palette

func register(p *palette.Palette) *palette.Palette {
	registry = append(registry, p)
	return p
}

// All returns every registered palette in registration order.
func All() []*palette.Palette {
	return registry
}
