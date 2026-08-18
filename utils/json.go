package utils

import (
	"encoding/json"
	"fmt"

	"github.com/leraniode/wondertone/palette"
	"github.com/leraniode/wondertone/tone"
)

// ToneJSON is a serialisable snapshot of a tone.
type ToneJSON struct {
	Name     string  `json:"name"`
	Hex      string  `json:"hex"`
	OKLCH    string  `json:"oklch"`
	Light    float64 `json:"light"`
	Vibrancy float64 `json:"vibrancy"`
	Hue      float64 `json:"hue"`
	Mood     string  `json:"mood,omitempty"`
	Energy   float64 `json:"energy,omitempty"`
}

// PaletteJSON is a serialisable snapshot of a palette.
type PaletteJSON struct {
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Tones       []ToneJSON `json:"tones"`
}

// ToToneJSON converts a tone to its JSON-friendly form.
func ToToneJSON(t tone.Tone) ToneJSON {
	return ToneJSON{
		Name:     t.Name(),
		Hex:      t.Hex(),
		OKLCH:    t.OKLCHString(),
		Light:    t.Light(),
		Vibrancy: t.Vibrancy(),
		Hue:      t.Hue(),
		Mood:     t.Mood(),
		Energy:   t.Energy(),
	}
}

// ToPaletteJSON converts a palette (name + tones) to JSON-friendly form.
// Description is left empty when exporting from a tone slice alone;
// pass the palette when available.
func ToPaletteJSON(name string, tones []tone.Tone) PaletteJSON {
	out := PaletteJSON{Name: name, Tones: make([]ToneJSON, 0, len(tones))}
	for _, t := range tones {
		out.Tones = append(out.Tones, ToToneJSON(t))
	}
	return out
}

// PaletteToJSON converts a *palette.Palette using its metadata.
func PaletteToJSON(p *palette.Palette) PaletteJSON {
	pj := ToPaletteJSON(p.Name(), p.All())
	pj.Description = p.Description()
	return pj
}

// JSON encodes tones as a JSON array of ToneJSON objects.
func JSON(tones []tone.Tone) ([]byte, error) {
	items := make([]ToneJSON, 0, len(tones))
	for _, t := range tones {
		items = append(items, ToToneJSON(t))
	}
	return json.MarshalIndent(items, "", "  ")
}

// JSONPalette encodes a palette as a JSON object.
func JSONPalette(p *palette.Palette) ([]byte, error) {
	return json.MarshalIndent(PaletteToJSON(p), "", "  ")
}

// JSONMust is like JSON but panics on error.
func JSONMust(tones []tone.Tone) string {
	b, err := JSON(tones)
	if err != nil {
		panic(fmt.Sprintf("utils: JSON encode: %v", err))
	}
	return string(b)
}
