package utils_test

import (
	"testing"

	"github.com/leraniode/glitter/tones"
	"github.com/leraniode/glitter/utils"
	"github.com/leraniode/wondertone/tone"
)

func TestByMood(t *testing.T) {
	quiet := utils.ByMood(tones.All(), "quiet")
	if len(quiet) == 0 {
		t.Fatal("expected at least one quiet-mood tone")
	}
	for _, tn := range quiet {
		if tn.Mood() != "quiet" {
			t.Errorf("expected mood quiet, got %q", tn.Mood())
		}
	}
}

func TestLightDark(t *testing.T) {
	all := tones.All()
	for _, tn := range utils.Light(all) {
		if !tn.IsLight() {
			t.Errorf("%q should be light", tn.Name())
		}
	}
	for _, tn := range utils.Dark(all) {
		if !tn.IsDark() {
			t.Errorf("%q should be dark", tn.Name())
		}
	}
}

func TestByHueRangeWrap(t *testing.T) {
	got := utils.ByHueRange(tones.All(), 350, 30)
	for _, tn := range got {
		h := tn.Hue()
		if !(h >= 350 || h <= 30) {
			t.Errorf("%q hue %v outside wrap range", tn.Name(), h)
		}
	}
}

func TestNames(t *testing.T) {
	all := tones.All()
	names := utils.Names(all)
	if len(names) != len(all) {
		t.Fatalf("Names len %d != %d", len(names), len(all))
	}
}

func TestFind(t *testing.T) {
	got, ok := utils.Find(tones.All(), func(tn tone.Tone) bool {
		return tn.Name() == "tone.red"
	})
	if !ok || got.Name() != "tone.red" {
		t.Errorf("Find tone.red failed")
	}
}
