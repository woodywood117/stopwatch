package stopwatch

import "time"

// Stopwatch is a simple stopwatch that provides functionality to start, pause, and reset.
type Stopwatch struct {
	start   time.Time
	paused  bool
	elapsed time.Duration
}

// New returns a new Stopwatch.
func New() *Stopwatch {
	return &Stopwatch{paused: true}
}

// Start starts the stopwatch.
func (s *Stopwatch) Start() {
	s.start = time.Now()
	s.paused = false
}

// Pause pauses the stopwatch.
func (s *Stopwatch) Pause() {
	if s.paused {
		return
	}
	s.elapsed += time.Since(s.start)
	s.paused = true
}

// Reset resets the stopwatch. Zeroing it out.
func (s *Stopwatch) Reset() {
	s.start = time.Now()
	s.paused = true
	s.elapsed = 0
}

// Elapsed returns the elapsed time.
func (s *Stopwatch) Elapsed() time.Duration {
	if s.paused {
		return s.elapsed
	}
	return s.elapsed + time.Since(s.start)
}
