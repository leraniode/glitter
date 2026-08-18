// Package utils provides helpers for querying, transforming, and exporting
// glitter tones and palettes — CSS variables, Tailwind themes, JSON, and more.
package utils

import "github.com/leraniode/wondertone/tone"

// Filter returns tones for which pred returns true, preserving order.
func Filter(tones []tone.Tone, pred func(tone.Tone) bool) []tone.Tone {
	out := make([]tone.Tone, 0, len(tones))
	for _, t := range tones {
		if pred(t) {
			out = append(out, t)
		}
	}
	return out
}

// ByMood returns tones whose Mood() equals mood (case-sensitive).
func ByMood(tones []tone.Tone, mood string) []tone.Tone {
	return Filter(tones, func(t tone.Tone) bool { return t.Mood() == mood })
}

// ByLightRange returns tones whose Light() is in [min, max] inclusive.
func ByLightRange(tones []tone.Tone, min, max float64) []tone.Tone {
	return Filter(tones, func(t tone.Tone) bool {
		l := t.Light()
		return l >= min && l <= max
	})
}

// ByHueRange returns tones whose Hue() is in [start, end].
// Handles wrap-around across 0° (e.g. start=350, end=20).
func ByHueRange(tones []tone.Tone, start, end float64) []tone.Tone {
	return Filter(tones, func(t tone.Tone) bool {
		h := t.Hue()
		if start <= end {
			return h >= start && h <= end
		}
		return h >= start || h <= end
	})
}

// Light returns only light tones (Light > 50).
func Light(tones []tone.Tone) []tone.Tone {
	return Filter(tones, tone.Tone.IsLight)
}

// Dark returns only dark tones (Light ≤ 50).
func Dark(tones []tone.Tone) []tone.Tone {
	return Filter(tones, tone.Tone.IsDark)
}

// Names returns the name of each tone in order.
func Names(tones []tone.Tone) []string {
	out := make([]string, len(tones))
	for i, t := range tones {
		out[i] = t.Name()
	}
	return out
}

// Map applies fn to every tone and returns the results in order.
func Map(tones []tone.Tone, fn func(tone.Tone) tone.Tone) []tone.Tone {
	out := make([]tone.Tone, len(tones))
	for i, t := range tones {
		out[i] = fn(t)
	}
	return out
}

// Find returns the first tone for which pred is true, and true.
func Find(tones []tone.Tone, pred func(tone.Tone) bool) (tone.Tone, bool) {
	for _, t := range tones {
		if pred(t) {
			return t, true
		}
	}
	return tone.Tone{}, false
}
