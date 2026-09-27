// Package palettes provides the official Leraniode color palettes built on
// wondertone.
//
// Available palettes:
//
//   - Leraniode — Leraniode full base palette (arc + accent + neutrals)
//   - Spectrum  — the identity arc itself: cyan through pink, one continuous line
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
