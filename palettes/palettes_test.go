package palettes_test

import (
	"testing"

	"github.com/leraniode/glitter/palettes"
)

func TestAll(t *testing.T) {
	if len(palettes.All()) == 0 {
		t.Fatal("expected registered palettes, got none")
	}
	names := map[string]bool{}
	for _, p := range palettes.All() {
		names[p.Name()] = true
	}
	for _, want := range []string{"spectrum", "leraniode"} {
		if !names[want] {
			t.Errorf("missing palette %q", want)
		}
	}
}

func TestLeraniodeContainsCores(t *testing.T) {
	p := palettes.Leraniode
	for _, name := range []string{"spectrum.violet", "spectrum.cyan", "confirmed", "neutral.page", "neutral.primary"} {
		if !p.Has(name) {
			t.Errorf("Leraniode missing tone %q", name)
		}
	}
}

func TestSpectrumSize(t *testing.T) {
	if n := len(palettes.Spectrum.All()); n != 5 {
		t.Errorf("Spectrum expected 5 tones, got %d", n)
	}
}

func TestGradientIsComputed(t *testing.T) {
	g := palettes.Gradient(9)
	if len(g) != 9 {
		t.Fatalf("expected 9 gradient stops, got %d", len(g))
	}
	if g[0].Hue() != 190 {
		t.Errorf("expected gradient to start at cyan (190°), got %.0f", g[0].Hue())
	}
	if g[len(g)-1].Hue() != 342 {
		t.Errorf("expected gradient to end at pink (342°), got %.0f", g[len(g)-1].Hue())
	}
	// Every stop should stay reasonably vivid — routed through the named
	// waypoints, not cut straight across the middle of the arc.
	for _, tn := range g {
		if tn.Vibrancy() < 60 {
			t.Errorf("stop at hue %.0f desaturated to vibrancy %.0f — gradient is cutting through the middle again", tn.Hue(), tn.Vibrancy())
		}
	}
}
