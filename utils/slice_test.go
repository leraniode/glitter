package utils_test

import (
	"testing"

	"github.com/leraniode/glitter/palettes"
	"github.com/leraniode/glitter/tones"
	"github.com/leraniode/glitter/utils"
)

func TestToneByName(t *testing.T) {
	all := tones.All()
	got, ok := utils.ToneByName(all, "tone.red")
	if !ok || got.Name() != "tone.red" {
		t.Fatalf("ToneByName tone.red failed")
	}
	if utils.HasTone(all, "nope") {
		t.Error("HasTone should be false for missing")
	}
}

func TestSliceHelpers(t *testing.T) {
	all := tones.All()
	n := len(all)
	if len(utils.First(all, 3)) > 3 {
		t.Error("First too long")
	}
	if len(utils.Last(all, 2)) > 2 {
		t.Error("Last too long")
	}
	if len(utils.Skip(all, n)) != 0 {
		t.Error("Skip all should be empty")
	}
	if len(utils.Range(all, 0, n)) != n {
		t.Error("Range full mismatch")
	}
}

func TestPaletteByName(t *testing.T) {
	all := palettes.All()
	p, ok := utils.PaletteByName(all, "spectrum")
	if !ok || p.Name() != "spectrum" {
		t.Fatal("PaletteByName spectrum failed")
	}
	names := utils.PaletteNames(all)
	if len(names) != len(all) {
		t.Fatal("PaletteNames length mismatch")
	}
}
