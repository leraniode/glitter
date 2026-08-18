package utils_test

import (
	"strings"
	"testing"

	"github.com/leraniode/glitter/palettes"
	"github.com/leraniode/glitter/tones"
	"github.com/leraniode/glitter/utils"
)

func TestMergeTones(t *testing.T) {
	a := tones.All()[:3]
	b := tones.All()[1:4] // overlap
	merged := utils.MergeTones(a, b)
	names := map[string]int{}
	for _, tn := range merged {
		names[tn.Name()]++
		if names[tn.Name()] > 1 {
			t.Errorf("duplicate %q after merge", tn.Name())
		}
	}
}

func TestGradientBetween(t *testing.T) {
	list := palettes.Leraniode.All()
	g, err := utils.GradientBetween(list, "niode", "lera", 5)
	if err != nil {
		t.Fatal(err)
	}
	if len(g) != 5 {
		t.Fatalf("expected 5 steps, got %d", len(g))
	}
}

func TestBuildPalette(t *testing.T) {
	subset := utils.ByMood(tones.All(), "warm")
	p, err := utils.BuildPalette("warm-accents", subset,
		utils.WithDescription("warm only"),
		utils.WithAuthor("test"),
	)
	if err != nil {
		t.Fatal(err)
	}
	if p.Name() != "warm-accents" {
		t.Errorf("name = %q", p.Name())
	}
	if len(p.All()) != len(subset) {
		t.Error("tone count mismatch")
	}
}

func TestGoSource(t *testing.T) {
	src := utils.GoSource(utils.First(tones.All(), 2), "demo")
	if !strings.Contains(src, "package demo") {
		t.Error("missing package")
	}
	if !strings.Contains(src, "tone.New") {
		t.Error("missing tone.New")
	}
}

func TestJSON(t *testing.T) {
	raw, err := utils.JSON(utils.First(tones.All(), 2))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"hex"`) {
		t.Error("expected hex field in JSON")
	}
}
