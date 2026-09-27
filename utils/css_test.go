package utils_test

import (
	"strings"
	"testing"

	"github.com/leraniode/glitter/palettes"
	"github.com/leraniode/glitter/utils"
)

func TestCSSVariables(t *testing.T) {
	css := utils.CSSVariables(palettes.Spectrum.All(), utils.CSSOptions{
		Prefix: "--spectrum-",
		Format: "hex",
	})
	if !strings.Contains(css, ":root") {
		t.Error("expected :root selector")
	}
	if !strings.Contains(css, "--spectrum-spectrum-violet:") {
		t.Error("expected --spectrum-spectrum-violet variable")
	}
	if !strings.Contains(css, "#") {
		t.Error("expected hex values")
	}
}

func TestCSSOKLCH(t *testing.T) {
	css := utils.CSSVariables(palettes.Spectrum.All(), utils.CSSOptions{
		Format: "oklch",
	})
	if !strings.Contains(css, "oklch(") {
		t.Error("expected oklch() values")
	}
}

func TestSCSSMap(t *testing.T) {
	scss := utils.SCSSMap(palettes.Spectrum.All(), "spectrum", utils.CSSOptions{})
	if !strings.HasPrefix(scss, "$spectrum:") {
		t.Error("expected $spectrum map")
	}
}

func TestTailwindTheme(t *testing.T) {
	tw := utils.TailwindTheme(palettes.Spectrum.All(), utils.CSSOptions{})
	if !strings.Contains(tw, "--color-") {
		t.Error("expected --color- prefix")
	}
}
