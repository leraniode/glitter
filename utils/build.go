package utils

import (
	"fmt"
	"strings"

	"github.com/leraniode/wondertone/mix"
	"github.com/leraniode/wondertone/palette"
	"github.com/leraniode/wondertone/tone"
)

// MergeTones concatenates multiple tone slices, skipping later duplicates by name.
func MergeTones(slices ...[]tone.Tone) []tone.Tone {
	seen := map[string]bool{}
	var out []tone.Tone
	for _, slice := range slices {
		for _, t := range slice {
			name := t.Name()
			if name != "" && seen[name] {
				continue
			}
			if name != "" {
				seen[name] = true
			}
			out = append(out, t)
		}
	}
	return out
}

// RenamePrefix returns a copy of tones with name prefix applied.
// "ember" + "accent." → "accent.ember". Empty names are left unchanged.
func RenamePrefix(tones []tone.Tone, prefix string) []tone.Tone {
	return Map(tones, func(t tone.Tone) tone.Tone {
		if t.Name() == "" {
			return t
		}
		return t.WithName(prefix + t.Name())
	})
}

// WithEnergyAll applies energy e to every tone.
func WithEnergyAll(tones []tone.Tone, e float64) []tone.Tone {
	return Map(tones, func(t tone.Tone) tone.Tone {
		return t.WithEnergy(e)
	})
}

// GradientBetween builds an n-step perceptual gradient between two named tones
// found in the provided slice. n must be >= 2.
func GradientBetween(tones []tone.Tone, from, to string, n int) ([]tone.Tone, error) {
	a, ok := ToneByName(tones, from)
	if !ok {
		return nil, fmt.Errorf("utils: gradient from tone %q not found", from)
	}
	b, ok := ToneByName(tones, to)
	if !ok {
		return nil, fmt.Errorf("utils: gradient to tone %q not found", to)
	}
	return mix.Gradient(a, b, n)
}

// MustGradientBetween is like GradientBetween but panics on error.
func MustGradientBetween(tones []tone.Tone, from, to string, n int) []tone.Tone {
	g, err := GradientBetween(tones, from, to, n)
	if err != nil {
		panic(err)
	}
	return g
}

// BuildPalette constructs a new palette from a name and tone list.
// Tones must already have unique non-empty names.
func BuildPalette(name string, tones []tone.Tone, opts ...PaletteOption) (*palette.Palette, error) {
	b := palette.New(name)
	cfg := paletteConfig{}
	for _, o := range opts {
		o(&cfg)
	}
	if cfg.description != "" {
		b = b.Description(cfg.description)
	}
	if cfg.author != "" {
		b = b.Author(cfg.author)
	}
	if cfg.version != "" {
		b = b.Version(cfg.version)
	}
	if cfg.mood != "" {
		b = b.Mood(cfg.mood)
	}
	for _, t := range tones {
		b = b.Add(t)
	}
	return b.Build()
}

// MustBuildPalette is like BuildPalette but panics on error.
func MustBuildPalette(name string, tones []tone.Tone, opts ...PaletteOption) *palette.Palette {
	p, err := BuildPalette(name, tones, opts...)
	if err != nil {
		panic(err)
	}
	return p
}

type paletteConfig struct {
	description, author, version, mood string
}

// PaletteOption configures BuildPalette.
type PaletteOption func(*paletteConfig)

// WithDescription sets the palette description.
func WithDescription(s string) PaletteOption {
	return func(c *paletteConfig) { c.description = s }
}

// WithAuthor sets the palette author.
func WithAuthor(s string) PaletteOption {
	return func(c *paletteConfig) { c.author = s }
}

// WithVersion sets the palette version.
func WithVersion(s string) PaletteOption {
	return func(c *paletteConfig) { c.version = s }
}

// WithMood sets the palette mood.
func WithMood(s string) PaletteOption {
	return func(c *paletteConfig) { c.mood = s }
}

// GoSource generates a rough Go source snippet that re-declares the tones
// as tone.New(...) calls. Useful for codegen / scaffolding.
func GoSource(tones []tone.Tone, packageName string) string {
	if packageName == "" {
		packageName = "palettes"
	}
	var b strings.Builder
	b.WriteString("package ")
	b.WriteString(packageName)
	b.WriteString("\n\n")
	b.WriteString("import \"github.com/leraniode/wondertone/tone\"\n\n")
	for _, t := range tones {
		name := t.Name()
		if name == "" {
			name = "unnamed"
		}
		varName := goIdent(name)
		b.WriteString(fmt.Sprintf("// %s — light=%.0f vibrancy=%.0f hue=%.0f mood=%q\n",
			name, t.Light(), t.Vibrancy(), t.Hue(), t.Mood()))
		b.WriteString(fmt.Sprintf("var %s = tone.New(\n", varName))
		b.WriteString(fmt.Sprintf("\ttone.Light(%.0f),\n", t.Light()))
		b.WriteString(fmt.Sprintf("\ttone.Vibrancy(%.0f),\n", t.Vibrancy()))
		b.WriteString(fmt.Sprintf("\ttone.Hue(%.0f),\n", t.Hue()))
		if t.Mood() != "" {
			b.WriteString(fmt.Sprintf("\ttone.Moody(%q),\n", t.Mood()))
		}
		b.WriteString(fmt.Sprintf("\ttone.Named(%q),\n", name))
		b.WriteString(")\n\n")
	}
	return b.String()
}

func goIdent(name string) string {
	parts := strings.FieldsFunc(name, func(r rune) bool {
		return r == '.' || r == '-' || r == '_' || r == ' '
	})
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "")
}
