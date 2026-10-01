package palettes_test

import (
	"testing"

	"github.com/leraniode/glitter/palettes"
)

func TestProductTonesExcludeOrgRange(t *testing.T) {
	for i, tn := range palettes.ProductTones {
		h := tn.Hue()
		if h > 280 && h < 360 {
			t.Errorf("ProductTones[%d] hue %.0f falls inside the org-reserved violet/pink range", i, h)
		}
	}
}

func TestProductAccentIndexesFixedSet(t *testing.T) {
	if n := len(palettes.ProductTones); n != 8 {
		t.Fatalf("expected 8 product tones, got %d", n)
	}
	if got, want := palettes.ProductAccent(0).Hue(), 20.0; got != want {
		t.Errorf("slot 0 should be red (20), got %.0f", got)
	}
	if got, want := palettes.ProductAccent(7).Hue(), 266.0; got != want {
		t.Errorf("slot 7 should be indigo (266), got %.0f", got)
	}
}

func TestGapTonesEvenlySpaced(t *testing.T) {
	// Red -> Amber -> Olive -> Moss -> Green should be four equal 33° steps
	// across the 132° gap, not a crowded leftover.
	hues := []float64{20, 53, 86, 119, 152}
	for i := 1; i < len(hues); i++ {
		if step := hues[i] - hues[i-1]; step != 33 {
			t.Errorf("step %d->%d is %.0f°, expected a uniform 33°", i-1, i, step)
		}
	}
}
