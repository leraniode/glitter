package utils

import (
	"github.com/leraniode/wondertone/palette"
	"github.com/leraniode/wondertone/tone"
)

// --- Tone slice helpers ---

// ToneByName returns the first tone with the given name.
func ToneByName(tones []tone.Tone, name string) (tone.Tone, bool) {
	for _, t := range tones {
		if t.Name() == name {
			return t, true
		}
	}
	return tone.Tone{}, false
}

// MustToneByName is like ToneByName but panics if not found.
func MustToneByName(tones []tone.Tone, name string) tone.Tone {
	t, ok := ToneByName(tones, name)
	if !ok {
		panic("utils: no tone named " + name)
	}
	return t
}

// HasTone reports whether a tone with the given name exists in the slice.
func HasTone(tones []tone.Tone, name string) bool {
	_, ok := ToneByName(tones, name)
	return ok
}

// First returns the first n tones. n is clamped to [0, len].
func First(tones []tone.Tone, n int) []tone.Tone {
	if n > len(tones) {
		n = len(tones)
	}
	if n < 0 {
		n = 0
	}
	return tones[:n]
}

// Last returns the last n tones. n is clamped to [0, len].
func Last(tones []tone.Tone, n int) []tone.Tone {
	if n > len(tones) {
		n = len(tones)
	}
	if n < 0 {
		n = 0
	}
	return tones[len(tones)-n:]
}

// Skip returns every tone after the first n. n is clamped to [0, len].
func Skip(tones []tone.Tone, n int) []tone.Tone {
	if n > len(tones) {
		n = len(tones)
	}
	if n < 0 {
		n = 0
	}
	return tones[n:]
}

// Range returns tones in the half-open interval [start, end).
func Range(tones []tone.Tone, start, end int) []tone.Tone {
	if start < 0 {
		start = 0
	}
	if end > len(tones) {
		end = len(tones)
	}
	if start > end {
		start = end
	}
	return tones[start:end]
}

// At returns the tone at 0-based index i.
func At(tones []tone.Tone, i int) (tone.Tone, bool) {
	if i < 0 || i >= len(tones) {
		return tone.Tone{}, false
	}
	return tones[i], true
}

// --- Palette slice helpers ---

// PaletteByName returns the first palette with the given name.
func PaletteByName(palettes []*palette.Palette, name string) (*palette.Palette, bool) {
	for _, p := range palettes {
		if p.Name() == name {
			return p, true
		}
	}
	return nil, false
}

// MustPaletteByName is like PaletteByName but panics if not found.
func MustPaletteByName(palettes []*palette.Palette, name string) *palette.Palette {
	p, ok := PaletteByName(palettes, name)
	if !ok {
		panic("utils: no palette named " + name)
	}
	return p
}

// HasPalette reports whether a palette with the given name exists.
func HasPalette(palettes []*palette.Palette, name string) bool {
	_, ok := PaletteByName(palettes, name)
	return ok
}

// PaletteNames returns the names of all palettes in order.
func PaletteNames(palettes []*palette.Palette) []string {
	names := make([]string, len(palettes))
	for i, p := range palettes {
		names[i] = p.Name()
	}
	return names
}
