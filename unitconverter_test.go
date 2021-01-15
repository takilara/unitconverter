package main

import (
	"testing"
	"math"
)

func TestFromSi(t *testing.T) {
	want := math.NaN()
	// want NaN on Nan value or "" string
	if got := FromSi(math.NaN(),"kg"); !math.IsNaN(got) {t.Errorf("FromSi() = %f, want %f", got, want)}
	if got := FromSi(33.3,""); !math.IsNaN(got) {t.Errorf("FromSi() = %f, want %f", got, want)}

}

func TestToSi(t *testing.T) {
        want := math.NaN()
        // want NaN on Nan value or "" string
        if got := ToSi(math.NaN(),"kg"); !math.IsNaN(got) {t.Errorf("ToSi() = %f, want %f", got, want)}
        if got := ToSi(33.3,""); !math.IsNaN(got) {t.Errorf("ToSi() = %f, want %f", got, want)}
}

