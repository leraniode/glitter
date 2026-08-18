package utils_test

import (
	"strings"
	"testing"

	"github.com/leraniode/glitter/palettes"
	"github.com/leraniode/glitter/utils"
)

func TestCSSVariables(t *testing.T) {
	css := utils.CSSVariables(palettes.LeraScale.All(), utils.CSSOptions{
		Prefix: "--lera-",
		Format: "hex",
	})
	if !strings.Contains(css, ":root") {
		t.Error("expected :root selector")
	}
	if !strings.Contains(css, "--lera-lera:") {
		t.Error("expected --lera-lera variable")
	}
	if !strings.Contains(css, "#") {
		t.Error("expected hex values")
	}
}

func TestCSSOKLCH(t *testing.T) {
	css := utils.CSSVariables(palettes.NiodeScale.All(), utils.CSSOptions{
		Format: "oklch",
	})
	if !strings.Contains(css, "oklch(") {
		t.Error("expected oklch() values")
	}
}

func TestSCSSMap(t *testing.T) {
	scss := utils.SCSSMap(palettes.Clarity.All(), "clarity", utils.CSSOptions{})
	if !strings.HasPrefix(scss, "$clarity:") {
		t.Error("expected $clarity map")
	}
}

func TestTailwindTheme(t *testing.T) {
	tw := utils.TailwindTheme(palettes.LeraScale.All(), utils.CSSOptions{})
	if !strings.Contains(tw, "--color-") {
		t.Error("expected --color- prefix")
	}
}
