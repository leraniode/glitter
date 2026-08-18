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
	for _, want := range []string{"lera", "niode", "clarity", "leraniode"} {
		if !names[want] {
			t.Errorf("missing palette %q", want)
		}
	}
}

func TestLeraniodeContainsCores(t *testing.T) {
	p := palettes.Leraniode
	for _, name := range []string{"lera", "niode", "clarity.sky", "ember", "ink", "abyss"} {
		if !p.Has(name) {
			t.Errorf("Leraniode missing tone %q", name)
		}
	}
}

func TestScalesSize(t *testing.T) {
	if n := len(palettes.LeraScale.All()); n != 5 {
		t.Errorf("LeraScale expected 5 tones, got %d", n)
	}
	if n := len(palettes.NiodeScale.All()); n != 5 {
		t.Errorf("NiodeScale expected 5 tones, got %d", n)
	}
	if n := len(palettes.Clarity.All()); n != 5 {
		t.Errorf("Clarity expected 5 tones, got %d", n)
	}
}
