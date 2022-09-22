package stopwatch

import "testing"

func TestNew(t *testing.T) {
	s := New()
	if s == nil {
		t.Error("New() returned nil")
	}
	if !s.paused {
		t.Error("New() returned a stopwatch that is not paused")
	}
	if s.elapsed != 0 {
		t.Error("New() returned a stopwatch with a non-zero elapsed time")
	}
}

func TestStart(t *testing.T) {
	s := New()
	s.Start()
	if s.paused {
		t.Error("Start() did not unpause a stopwatch")
	}
}

func TestPause(t *testing.T) {
	s := New()
	s.Start()
	s.Pause()
	if !s.paused {
		t.Error("Pause() did not pause a stopwatch")
	}
}

func TestReset(t *testing.T) {
	s := New()
	s.Start()
	s.Reset()
	if !s.paused {
		t.Error("Reset() did not pause a stopwatch")
	}
	if s.elapsed != 0 {
		t.Error("Reset() did not reset the elapsed time")
	}
}

func TestElapsed(t *testing.T) {
	s := New()
	s.Start()
	s.Pause()
	if s.Elapsed() != s.elapsed {
		t.Error("Elapsed() did not return the elapsed time")
	}
}
