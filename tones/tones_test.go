package tones_test

import (
	"testing"

	"github.com/leraniode/glitter/tones"
)

func TestAllNotEmpty(t *testing.T) {
	if len(tones.All()) == 0 {
		t.Fatal("expected registered tones, got none")
	}
}

func TestKnownTonesPresent(t *testing.T) {
	names := map[string]bool{}
	for _, tn := range tones.All() {
		names[tn.Name()] = true
	}
	for _, want := range []string{"neutral.page", "neutral.primary", "tone.green", "tone.red"} {
		if !names[want] {
			t.Errorf("expected tone %q to be registered", want)
		}
	}
}
