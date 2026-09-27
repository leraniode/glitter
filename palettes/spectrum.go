package palettes

import (
	"github.com/leraniode/glitter/utils"
	"github.com/leraniode/wondertone/palette"
	"github.com/leraniode/wondertone/tone"
)

// Spectrum is the Leraniode identity arc: one continuous line from cyan
// through blue, indigo, and violet, to pink — not five separate named
// families. Every tone shares the exact same recipe (Light 64, Vibrancy 92);
// only Hue changes, in a clean 38° step. The relationship between any two
// tones here is arithmetic, not chosen by eye — that's the whole point.
//
// Anything that isn't part of the arc's own math doesn't belong in it.
// See tones.Confirmed for the one deliberate exception.

var spectrumCyan = tone.New(
	tone.Light(64),
	tone.Vibrancy(92),
	tone.Hue(190),
	tone.Moody("clear"),
	tone.Named("spectrum.cyan"),
)

var spectrumBlue = tone.New(
	tone.Light(64),
	tone.Vibrancy(92),
	tone.Hue(228),
	tone.Moody("clear"),
	tone.Named("spectrum.blue"),
)

var spectrumIndigo = tone.New(
	tone.Light(64),
	tone.Vibrancy(92),
	tone.Hue(266),
	tone.Moody("clear"),
	tone.Named("spectrum.indigo"),
)

// spectrumViolet sits at hue 304 — deliberately close to the retired niode
// core (297°). The one thread that survived a full teardown, kept on
// purpose this time instead of by accident.
var spectrumViolet = tone.New(
	tone.Light(64),
	tone.Vibrancy(92),
	tone.Hue(304),
	tone.Moody("clear"),
	tone.Named("spectrum.violet"),
)

var spectrumPink = tone.New(
	tone.Light(64),
	tone.Vibrancy(92),
	tone.Hue(342),
	tone.Moody("clear"),
	tone.Named("spectrum.pink"),
)

// Spectrum is the full five-stop identity arc, cyan to pink.
var Spectrum = register(palette.New("spectrum").
	Description("The Leraniode identity arc. One continuous line, cyan to pink.").
	Author("Leraniode").
	Version("0.1.0").
	Mood("clear").
	Add(spectrumCyan).
	Add(spectrumBlue).
	Add(spectrumIndigo).
	Add(spectrumViolet).
	Add(spectrumPink).
	MustBuild())

// arcOrder is the sequence of named waypoints Gradient walks, in order.
var arcOrder = []string{
	"spectrum.cyan", "spectrum.blue", "spectrum.indigo", "spectrum.violet", "spectrum.pink",
}

// Gradient returns n perceptually interpolated stops across the full arc,
// cyan to pink, routed *through* the four named waypoints in between —
// not a single straight hop from one end to the other.
//
// A direct two-point mix between hues this far apart cuts a chord across
// OKLab space and desaturates badly in the middle; walking the arc segment
// by segment (cyan→blue→indigo→violet→pink) keeps every stop as vivid as
// its nearest named anchor. Use this instead of hand-picking intermediate
// colors — any point on the arc is computed, never eyeballed.
func Gradient(n int) []tone.Tone {
	if n < len(arcOrder) {
		n = len(arcOrder)
	}
	segments := len(arcOrder) - 1
	totalSteps := n - 1
	base, rem := totalSteps/segments, totalSteps%segments

	all := Spectrum.All()
	out := make([]tone.Tone, 0, n)
	for i := 0; i < segments; i++ {
		steps := base
		if i < rem {
			steps++
		}
		segN := steps + 1
		if segN < 2 {
			segN = 2
		}
		seg := utils.MustGradientBetween(all, arcOrder[i], arcOrder[i+1], segN)
		if i > 0 {
			seg = seg[1:] // drop the point shared with the previous segment
		}
		out = append(out, seg...)
	}
	return out
}
