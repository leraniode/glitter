package utils

import (
	"fmt"
	"strings"

	"github.com/leraniode/wondertone/tone"
)

// CSSOptions controls how CSS custom properties are generated.
type CSSOptions struct {
	// Selector is the CSS rule selector. Default: ":root".
	Selector string
	// Prefix is prepended to every variable name, e.g. "--color".
	// Default: "--".
	Prefix string
	// Format selects the colour function written as the value.
	// One of: "hex" (default), "oklch", "rgb".
	Format string
	// Indent is the indentation string for each declaration. Default: "  ".
	Indent string
}

func (o CSSOptions) withDefaults() CSSOptions {
	if o.Selector == "" {
		o.Selector = ":root"
	}
	if o.Prefix == "" {
		o.Prefix = "--"
	}
	if o.Format == "" {
		o.Format = "hex"
	}
	if o.Indent == "" {
		o.Indent = "  "
	}
	return o
}

// cssName turns a tone name into a safe CSS custom-property suffix.
// "lera.petal" → "lera-petal", "ink" → "ink".
func cssName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "_", "-")
	name = strings.ReplaceAll(name, ".", "-")
	return name
}

func formatColour(t tone.Tone, format string) string {
	switch strings.ToLower(format) {
	case "oklch":
		return "oklch(" + t.OKLCHString() + ")"
	case "rgb":
		r, g, b := t.RGB()
		return fmt.Sprintf("rgb(%d %d %d)", r, g, b)
	default: // hex
		return t.Hex()
	}
}

// CSSVariables renders tones as a CSS custom-properties block.
//
//	:root {
//	  --lera: #f24986;
//	  --niode: #9049fb;
//	}
func CSSVariables(tones []tone.Tone, opts CSSOptions) string {
	opts = opts.withDefaults()
	var b strings.Builder
	b.WriteString(opts.Selector)
	b.WriteString(" {\n")
	for _, t := range tones {
		name := t.Name()
		if name == "" {
			continue
		}
		b.WriteString(opts.Indent)
		b.WriteString(opts.Prefix)
		b.WriteString(cssName(name))
		b.WriteString(": ")
		b.WriteString(formatColour(t, opts.Format))
		b.WriteString(";\n")
	}
	b.WriteString("}\n")
	return b.String()
}

// CSSVariablesFlat returns only the declarations (no selector wrapper),
// useful for injecting into an existing rule.
func CSSVariablesFlat(tones []tone.Tone, opts CSSOptions) string {
	opts = opts.withDefaults()
	var b strings.Builder
	for _, t := range tones {
		name := t.Name()
		if name == "" {
			continue
		}
		b.WriteString(opts.Prefix)
		b.WriteString(cssName(name))
		b.WriteString(": ")
		b.WriteString(formatColour(t, opts.Format))
		b.WriteString(";\n")
	}
	return b.String()
}

// SCSSMap renders tones as an SCSS map.
//
//	$leraniode: (
//	  "lera": #f24986,
//	  "niode": #9049fb,
//	);
func SCSSMap(tones []tone.Tone, mapName string, opts CSSOptions) string {
	opts = opts.withDefaults()
	if mapName == "" {
		mapName = "colors"
	}
	var b strings.Builder
	b.WriteString("$")
	b.WriteString(mapName)
	b.WriteString(": (\n")
	for i, t := range tones {
		name := t.Name()
		if name == "" {
			continue
		}
		b.WriteString(opts.Indent)
		b.WriteString(`"`)
		b.WriteString(cssName(name))
		b.WriteString(`": `)
		b.WriteString(formatColour(t, opts.Format))
		if i < len(tones)-1 {
			b.WriteString(",")
		}
		b.WriteString("\n")
	}
	b.WriteString(");\n")
	return b.String()
}

// TailwindTheme renders a Tailwind CSS v3/v4 theme colour fragment
// suitable for pasting into tailwind.config or @theme.
//
//	--color-lera: #f24986;
//	--color-niode: #9049fb;
func TailwindTheme(tones []tone.Tone, opts CSSOptions) string {
	opts = opts.withDefaults()
	if opts.Prefix == "--" {
		opts.Prefix = "--color-"
	}
	return CSSVariablesFlat(tones, opts)
}
