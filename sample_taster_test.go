package sample_taster

import "testing"

func TestTaste(t *testing.T) {
    want := "The sample tastes like: 'This is a great sample!'"
    if got := Taste(); got != want {
        t.Errorf("Taste() = %q, want %q", got, want)
    }
}
